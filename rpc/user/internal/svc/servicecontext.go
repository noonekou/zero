package svc

import (
	"bookstore/common/model"
	"bookstore/rpc/user/internal/config"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	Conn               sqlx.SqlConn
	UserModel          model.TUserModel // 用户模型
	AdminUserModel     model.TAdminUserModel
	AdminUserRoleModel model.TAdminUserRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:             c,
		Conn:               conn,
		UserModel:          model.NewTUserModel(conn),
		AdminUserModel:     model.NewTAdminUserModel(conn),
		AdminUserRoleModel: model.NewTAdminUserRoleModel(conn),
	}
}
