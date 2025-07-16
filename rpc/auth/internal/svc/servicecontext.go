package svc

import (
	"bookstore/rpc/auth/internal/config"
	"bookstore/rpc/model"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserModel           model.TUserModel
	AdminUserModel      model.TAdminUserModel
	RoleModel           model.TRoleModel
	PermissionModel     model.TPermissionModel
	UserRoleModel       model.TUserRoleModel
	RolePermissionModel model.TRolePermissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:              c,
		UserModel:           model.NewTUserModel(conn),
		AdminUserModel:      model.NewTAdminUserModel(conn),
		RoleModel:           model.NewTRoleModel(conn),
		PermissionModel:     model.NewTPermissionModel(conn),
		UserRoleModel:       model.NewTUserRoleModel(conn),
		RolePermissionModel: model.NewTRolePermissionModel(conn),
	}
}
