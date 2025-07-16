package user

import (
	"net/http"

	"bookstore/api/internal/logic/user"
	"bookstore/api/internal/svc"
	"bookstore/response"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		response.Response(w, resp, err)
	}
}
