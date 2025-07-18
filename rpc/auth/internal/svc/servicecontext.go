package svc

import (
	"bookstore/common/model"
	"bookstore/rpc/auth/internal/config"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserModel           model.TUserModel
	AdminUserModel      model.TAdminUserModel
	PermissionModel     model.TPermissionModel
	ResourceModel       model.TResourceModel
	RoleModel           model.TRoleModel
	RolePermissionModel model.TRolePermissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:              c,
		UserModel:           model.NewTUserModel(conn),
		AdminUserModel:      model.NewTAdminUserModel(conn),
		PermissionModel:     model.NewTPermissionModel(conn),
		ResourceModel:       model.NewTResourceModel(conn),
		RoleModel:           model.NewTRoleModel(conn),
		RolePermissionModel: model.NewTRolePermissionModel(conn),
	}
}
