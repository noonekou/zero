package auth

import (
	"net/http"

	"bookstore/admin/internal/logic/auth"
	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Role
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewAddRoleLogic(r.Context(), svcCtx)
		resp, err := l.AddRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
