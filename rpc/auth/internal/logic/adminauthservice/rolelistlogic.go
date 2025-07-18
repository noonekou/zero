package adminauthservicelogic

import (
	"context"

	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *auth.PageReq) (*auth.RoleListResp, error) {
	total, err := l.svcCtx.RoleModel.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	roles, err := l.svcCtx.RoleModel.FindByPage(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*auth.Role, 0)
	for _, v := range *roles {
		children := make([]*auth.Permission, 0)
		permissions, err := l.svcCtx.RolePermissionModel.FindByRoleName(l.ctx, v.Name)
		if err != nil {
			continue
		}

		for _, p := range permissions {
			children = append(children, &auth.Permission{Id: p.Id, Code: 0, Description: p.PermissionName, ParentCode: 0, Children: nil, CreatedAt: v.CreatedAt.Unix(), UpdatedAt: v.UpdatedAt.Unix()})
		}

		list = append(list, &auth.Role{Id: v.Id, Name: v.Name, Permissions: children, CreatedAt: v.CreatedAt.Unix(), UpdatedAt: v.UpdatedAt.Unix()})
	}

	return &auth.RoleListResp{
		Total: total,
		List:  list,
	}, nil
}
