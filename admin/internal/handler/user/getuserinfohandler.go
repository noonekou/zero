package user

import (
	"net/http"

	"bookstore/admin/internal/logic/user"
	"bookstore/admin/internal/svc"
	"bookstore/common/response"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		response.Response(w, resp, err)
	}
}
