package user

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	errs "bookstore/common/error"
	"bookstore/rpc/user/client/adminuserservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	user, err := l.svcCtx.AdminUser.GetUserInfo(l.ctx, &adminuserservice.GetUserInfoReq{Id: l.ctx.Value(types.CtxKeyUserID).(int64)})
	if err != nil {
		return nil, err
	}

	if user.Info == nil {
		return nil, errs.ErrUserNotFound
	}

	return &types.GetUserInfoResp{UserInfo: types.UserInfo{Id: user.Info.Id, UserName: user.Info.UserName, NickName: user.Info.NickName, Avatar: user.Info.Avatar, Email: user.Info.Email, Phone: user.Info.Phone, Status: int(user.Info.Status), CreatedAt: user.Info.CreatedAt, UpdatedAt: user.Info.UpdatedAt}}, nil
}
