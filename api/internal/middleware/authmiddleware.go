package middleware

import (
	"bookstore/api/internal/pkg/auth"
	"bookstore/api/internal/types"
	"bookstore/response"
	"context"
	"net/http"
)

type AuthMiddleware struct {
	AccessSecret string
	AccessExpire int64
}

func NewAuthMiddleware(accessSecret string, accessExpire int64) *AuthMiddleware {
	return &AuthMiddleware{
		AccessSecret: accessSecret,
		AccessExpire: accessExpire,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId, err := auth.ValidateToken(m.AccessSecret, token)

		if err != nil {
			response.ResponseError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), types.CtxKeyUserID, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
