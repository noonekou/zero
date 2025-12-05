package adminauthservicelogic

import (
	"context"
	"database/sql"

	errs "bookstore/common/error"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	if in.Id == 0 {
		return nil, errs.ErrRoleNotFound.GRPCStatus().Err()
	}

	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if role == nil {
		return nil, errs.ErrRoleNotFound.GRPCStatus().Err()
	}

	// 检查角色是否有关联用户
	userCount, err := l.svcCtx.AdminUserRoleModel.CountByRoleId(l.ctx, in.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if userCount > 0 {
		return nil, errs.ErrRoleHasUsers.GRPCStatus().Err()
	}

	err = l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// Use session-based models to ensure operations run within transaction
		roleModel := l.svcCtx.RoleModel.WithSession(session)
		rolePermissionModel := l.svcCtx.RolePermissionModel.WithSession(session)

		err = roleModel.Delete(ctx, in.Id)
		if err != nil {
			return err
		}
		err = rolePermissionModel.DeleteByRoleName(ctx, role.Name)
		return err
	})

	return &auth.Empty{}, err
}
