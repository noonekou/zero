// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package permission

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/auth/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 权限列表
func NewPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionListLogic {
	return &PermissionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionListLogic) PermissionList() (resp []types.Permission, err error) {
	permissionList, err := l.svcCtx.Auth.PermissionList(l.ctx, &auth.PermissionListReq{})
	if err != nil {
		return nil, err
	}

	permissions := make([]types.Permission, 0)
	for _, v := range permissionList.List {
		children := make([]types.Permission, 0)
		for _, c := range v.Children {
			children = append(children, types.Permission{Id: c.Id, Code: int(c.Code), Description: c.Description, ParentCode: int(c.ParentCode), Children: make([]types.Permission, 0), CreatedAt: c.CreatedAt, UpdatedAt: c.UpdatedAt})
		}
		permissions = append(permissions, types.Permission{Id: v.Id, Code: int(v.Code), Description: v.Description, ParentCode: int(v.ParentCode), Children: children, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt})
	}

	return permissions, nil
}
