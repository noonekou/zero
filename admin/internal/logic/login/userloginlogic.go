package login

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	rsp, err := l.svcCtx.Auth.Login(l.ctx, &auth.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Token: rsp.Token,
		User: types.UserInfo{
			Id:        user.Info.Id,
			UserName:  user.Info.UserName,
			NickName:  user.Info.NickName,
			Avatar:    user.Info.Avatar,
			Email:     user.Info.Email,
			Phone:     user.Info.Phone,
			Status:    int(user.Info.Status),
			CreatedAt: user.Info.CreatedAt,
			UpdatedAt: user.Info.UpdatedAt,
		},
	}, nil
}
