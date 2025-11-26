package auth

import (
	"net/http"

	"bookstore/admin/internal/logic/auth"
	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	"bookstore/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewAuthLoginLogic(r.Context(), svcCtx)
		resp, err := l.AuthLogin(&req)
		response.Response(w, resp, err)
	}
}
