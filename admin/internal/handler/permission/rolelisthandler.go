// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package permission

import (
	"net/http"

	"bookstore/admin/internal/logic/permission"
	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 角色列表
func RoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permission.NewRoleListLogic(r.Context(), svcCtx)
		resp, err := l.RoleList(&req)
		response.Response(w, resp, err)
	}
}
