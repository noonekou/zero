package auth

import (
	"net/http"

	"bookstore/api/internal/logic/auth"
	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewAuthRegisterLogic(r.Context(), svcCtx)
		resp, err := l.AuthRegister(&req)
		response.Response(w, resp, err)
	}
}
