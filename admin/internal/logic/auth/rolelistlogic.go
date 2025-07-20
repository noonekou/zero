package auth

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/rpc/auth/auth"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.PageReq) (resp *types.RoleListResp, err error) {
	roleList, err := l.svcCtx.Auth.RoleList(l.ctx, &auth.PageReq{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}

	resp = &types.RoleListResp{
		Total: roleList.Total,
		List: lo.Map(roleList.List, func(item *auth.Role, _ int) types.Role {
			return types.Role{Id: item.Id, Name: item.Name, Permissions: lo.Map(item.Permissions, func(item *auth.Permission, _ int) types.Permission {
				return types.Permission{Id: item.Id, Code: int(item.Code), Description: item.Description, ParentCode: int(item.ParentCode)}
			}), CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt}
		}),
	}

	return resp, nil
}
