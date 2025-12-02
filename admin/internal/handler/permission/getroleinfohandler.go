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

// 获取角色信息
func GetRoleInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permission.NewGetRoleInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
