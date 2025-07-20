package auth

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/auth/auth"

	"github.com/samber/lo"
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
	role, error := l.svcCtx.Auth.GetRoleInfo(l.ctx, &auth.RoleInfoReq{Id: req.Id})
	if error != nil {
		return nil, error
	}

	return &types.Role{
		Id:   role.Id,
		Name: role.Name,
		Permissions: lo.Map(role.Permissions, func(item *auth.Permission, _ int) types.Permission {
			return types.Permission{Id: item.Id, Code: int(item.Code), Description: item.Description, ParentCode: int(item.ParentCode)}
		}),
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}, error
}
