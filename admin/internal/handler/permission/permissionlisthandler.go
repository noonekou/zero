// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package permission

import (
	"net/http"

	"bookstore/admin/internal/logic/permission"
	"bookstore/admin/internal/svc"
	"bookstore/common/response"
)

// 权限列表
func PermissionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := permission.NewPermissionListLogic(r.Context(), svcCtx)
		resp, err := l.PermissionList()
		response.Response(w, resp, err)
	}
}
