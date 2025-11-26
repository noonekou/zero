package auth

import (
	"net/http"

	"bookstore/admin/internal/logic/auth"
	"bookstore/admin/internal/svc"
	"bookstore/common/response"
)

func PermissionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewPermissionListLogic(r.Context(), svcCtx)
		resp, err := l.PermissionList()
		response.Response(w, resp, err)
	}
}
