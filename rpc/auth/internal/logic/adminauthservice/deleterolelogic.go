package adminauthservicelogic

import (
	"context"
	"database/sql"

	errs "bookstore/common/error"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *auth.RoleInfoReq) (*auth.Empty, error) {
	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if role == nil {
		return nil, errs.ErrRoleNotFound.GRPCStatus().Err()
	}

	err = l.svcCtx.RoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RolePermissionModel.DeleteByRoleName(l.ctx, role.Name)
	if err != nil {
		return nil, err
	}

	return &auth.Empty{}, nil
}
