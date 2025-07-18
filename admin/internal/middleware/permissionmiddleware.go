package middleware

import (
	"bookstore/admin/internal/types"
	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/response"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type PermissionMiddleware struct {
	dbConn sqlx.SqlConn
}

func NewPermissionMiddleware(dbConn sqlx.SqlConn) *PermissionMiddleware {
	return &PermissionMiddleware{
		dbConn: dbConn,
	}
}

func (m *PermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiPermissionModel := model.NewTApiPermissionModel(m.dbConn)
		rolePermissionModel := model.NewTRolePermissionModel(m.dbConn)

		method := strings.ToUpper(r.Method)
		path := r.URL.Path

		permissionName, err := apiPermissionModel.FindOneByMethodAndPath(r.Context(), method, path)
		if err != nil {
			response.ResponseError(w, errs.ErrNoPermission)
			return
		}

		if permissionName == "*" {
			next(w, r)
			return
		}

		userId := r.Context().Value(types.CtxKeyUserID).(int64)
		permissionNames, err := rolePermissionModel.FindPermissionNameByUserId(r.Context(), userId)
		if err != nil {
			response.ResponseError(w, errs.ErrNoPermission)
			return
		}

		for _, v := range permissionNames {
			if v == permissionName {
				next(w, r)
				return
			}
		}

		response.ResponseError(w, errs.ErrNoPermission)
	}
}
