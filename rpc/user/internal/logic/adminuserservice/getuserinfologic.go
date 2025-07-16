package adminuserservicelogic

import (
	"context"
	"database/sql"

	errs "bookstore/common/error"
	"bookstore/rpc/user/internal/svc"
	"bookstore/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	tUser, err := l.svcCtx.AdminUserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if tUser == nil {
		return nil, errs.ErrUserNotFound.GRPCStatus().Err()
	}

	return &user.GetUserInfoResp{Info: &user.UserInfo{Id: tUser.Id, UserName: tUser.Username, NickName: tUser.Nickname, Avatar: tUser.Avatar, Email: tUser.Email, Phone: tUser.Phone, Status: int32(tUser.Status), CreatedAt: tUser.CreatedAt.Unix(), UpdatedAt: tUser.UpdatedAt.Unix()}}, nil
}
