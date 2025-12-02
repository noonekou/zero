// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package permission

import (
	"net/http"

	"bookstore/admin/internal/logic/permission"
	"bookstore/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 权限列表
func PermissionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := permission.NewPermissionListLogic(r.Context(), svcCtx)
		resp, err := l.PermissionList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
