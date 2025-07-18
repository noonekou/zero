package adminauthservicelogic

import (
	"context"
	"database/sql"

	errs "bookstore/common/error"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleInfoLogic {
	return &GetRoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleInfoLogic) GetRoleInfo(in *auth.RoleInfoReq) (*auth.Role, error) {
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if role == nil {
		return nil, errs.ErrRoleNotFound.GRPCStatus().Err()
	}

	permissions, err := l.svcCtx.RolePermissionModel.FindByRoleName(l.ctx, role.Name)
	if err != nil {
		return nil, err
	}

	children := make([]*auth.Permission, 0)
	for _, p := range permissions {
		children = append(children, &auth.Permission{Id: p.Id, Code: 0, Description: p.PermissionName, ParentCode: 0, Children: nil, CreatedAt: role.CreatedAt.Unix(), UpdatedAt: role.UpdatedAt.Unix()})
	}

	return &auth.Role{
		Id:          role.Id,
		Name:        role.Name,
		Permissions: children,
		CreatedAt:   role.CreatedAt.Unix(),
		UpdatedAt:   role.UpdatedAt.Unix(),
	}, nil
}
