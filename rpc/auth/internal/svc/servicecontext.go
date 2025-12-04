package svc

import (
	"bookstore/common/model"
	"bookstore/rpc/auth/internal/config"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	UserModel           model.TUserModel
	AdminUserModel      model.TAdminUserModel
	AdminUserRoleModel  model.TAdminUserRoleModel
	PermissionModel     model.TPermissionModel
	ResourceModel       model.TResourceModel
	RoleModel           model.TRoleModel
	RolePermissionModel model.TRolePermissionModel
	Conn                sqlx.SqlConn
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:              c,
		Conn:                conn,
		UserModel:           model.NewTUserModel(conn),
		AdminUserModel:      model.NewTAdminUserModel(conn),
		AdminUserRoleModel:  model.NewTAdminUserRoleModel(conn),
		PermissionModel:     model.NewTPermissionModel(conn),
		ResourceModel:       model.NewTResourceModel(conn),
		RoleModel:           model.NewTRoleModel(conn),
		RolePermissionModel: model.NewTRolePermissionModel(conn),
		RedisClient:         redis.MustNewRedis(c.TokenRedis),
	}
}
