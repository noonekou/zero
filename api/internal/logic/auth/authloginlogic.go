package auth

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	errs "bookstore/common/error"
	"bookstore/rpc/auth/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLoginLogic {
	return &AuthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLoginLogic) AuthLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errs.ErrUsernameOrPasswordIsEmpty
	}

	rsp, err := l.svcCtx.Auth.Login(l.ctx, &auth.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Token: rsp.Token,
	}, nil
}
