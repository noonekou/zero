// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package permission

import (
	"net/http"

	"bookstore/admin/internal/logic/permission"
	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加角色
func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Role
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permission.NewAddRoleLogic(r.Context(), svcCtx)
		err := l.AddRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
