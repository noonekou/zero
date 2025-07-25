// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.5

package handler

import (
	"net/http"

	auth "bookstore/admin/internal/handler/auth"
	user "bookstore/admin/internal/handler/user"
	"bookstore/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/auth/login",
				Handler: auth.AuthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/auth/permission/list",
				Handler: auth.PermissionListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/register",
				Handler: auth.AuthRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/role/add",
				Handler: auth.AddRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/auth/role/delete",
				Handler: auth.DeleteRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/auth/role/info",
				Handler: auth.GetRoleInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/auth/role/list",
				Handler: auth.RoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/auth/role/update",
				Handler: auth.UpdateRoleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware, serverCtx.PermissionMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/list",
					Handler: user.UserListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
	)
}
