package auth

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/auth/auth"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.Role) (err error) {
	_, error := l.svcCtx.Auth.UpdateRole(l.ctx, &auth.Role{
		Id:     req.Id,
		Name:   req.Name,
		Status: int64(req.Status),
		Permissions: lo.Map(req.Permissions, func(item types.Permission, _ int) *auth.Permission {
			return &auth.Permission{Id: item.Id, Code: int32(item.Code), Description: item.Description, ParentCode: int32(item.ParentCode)}
		}),
	})

	return error
}
