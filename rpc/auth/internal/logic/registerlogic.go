package logic

import (
	"context"
	"database/sql"

	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"
	"bookstore/rpc/model"

	common "bookstore/common/auth"
	errs "bookstore/common/error"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *auth.RegisterReq) (*auth.RegisterResp, error) {
	if (in.Username == "" || in.Password == "") && in.ConfirmPassword == "" {
		return nil, errs.ErrUsernameOrPasswordIsEmpty.GRPCStatus().Err()
	}
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if user != nil {
		return nil, errs.ErrUsernameAlreadyExist.GRPCStatus().Err()
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.TAdminUser{
		Username: in.Username,
		Password: in.Password,
		Nickname: "",
		Avatar:   "",
		Email:    "",
		Phone:    "",
		Status:   1,
	})

	if err != nil {
		return nil, err
	}

	user, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	token, err := common.GenerateToken(l.svcCtx.Config.Authorization.AccessSecret, l.svcCtx.Config.Authorization.AccessExpire, user.Id)

	if err != nil {
		return nil, err
	}

	return &auth.RegisterResp{
		Token: token,
	}, nil
}
