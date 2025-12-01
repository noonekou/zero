package adminuserservicelogic

import (
	"context"

	errs "bookstore/common/error"
	"bookstore/rpc/user/internal/svc"
	"bookstore/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.UserInfoReq) (*user.Empty, error) {
	if in.Id == 0 {
		return nil, errs.ErrUserNotFound.GRPCStatus().Err()
	}

	err := l.svcCtx.AdminUserModel.Delete(l.ctx, in.Id)
	return &user.Empty{}, err
}
