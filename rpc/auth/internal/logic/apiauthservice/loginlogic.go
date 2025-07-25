package apiauthservicelogic

import (
	common "bookstore/common/auth"
	errs "bookstore/common/error"

	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"
	"context"
	"database/sql"

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
		return nil, errs.ErrUsernameOrPasswordIsEmpty.GRPCStatus().Err()
	}

	tUser, err := l.svcCtx.UserModel.FindOneByUsernameAndPassword(l.ctx, in.Username, in.Password)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if tUser == nil {
		return nil, errs.ErrUsernameNotExist.GRPCStatus().Err()
	}

	token, err := common.GenerateToken(l.svcCtx.Config.Authorization.AccessSecret, l.svcCtx.Config.Authorization.AccessExpire, tUser.Id)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResp{Token: token}, nil
}
