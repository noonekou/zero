package auth

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleInfoLogic {
	return &GetRoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleInfoLogic) GetRoleInfo(req *types.RoleInfoReq) (resp *types.Role, err error) {
	// todo: add your logic here and delete this line

	return
}
