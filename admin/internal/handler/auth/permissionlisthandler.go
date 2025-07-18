package auth

import (
	"net/http"

	"bookstore/admin/internal/logic/auth"
	"bookstore/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PermissionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewPermissionListLogic(r.Context(), svcCtx)
		resp, err := l.PermissionList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
