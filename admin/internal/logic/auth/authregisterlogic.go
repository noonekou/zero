package auth

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/auth/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthRegisterLogic {
	return &AuthRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthRegisterLogic) AuthRegister(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if (req.Username == "" || req.Password == "") && req.Confirm_password == "" {
		return nil, err
	}

	rsp, err := l.svcCtx.Auth.Register(l.ctx, &auth.RegisterReq{
		Username:        req.Username,
		Password:        req.Password,
		ConfirmPassword: req.Confirm_password,
	})

	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Token: rsp.Token,
	}, nil
}
