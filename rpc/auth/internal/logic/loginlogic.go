package logic

import (
	"context"
	"database/sql"
	"errors"

	common "bookstore/common/auth"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *auth.LoginReq) (*auth.LoginResp, error) {
	if in.Username == "" || in.Password == "" {
		return nil, errors.New("username or password is empty")
	}

	tUser, err := l.svcCtx.UserModel.FindOneByUsernameAndPassword(l.ctx, in.Username, in.Password)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if tUser == nil {
		return nil, errors.New("username not exist")
	}

	token, err := common.GenerateToken(l.svcCtx.Config.Authorization.AccessSecret, l.svcCtx.Config.Authorization.AccessExpire, tUser.Id)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResp{Token: token}, nil
}
