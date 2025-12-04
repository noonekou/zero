package middleware

import (
	"bookstore/admin/internal/types"
	"bookstore/common/auth"
	"bookstore/response"
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type AuthMiddleware struct {
	AccessSecret string
	AccessExpire int64
	RedisClient  *redis.Redis
}

func NewAuthMiddleware(accessSecret string, accessExpire int64, redisClient *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{
		AccessSecret: accessSecret,
		AccessExpire: accessExpire,
		RedisClient:  redisClient,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// 1. Validate JWT signature and format
		userId, err := auth.ValidateToken(m.AccessSecret, token)
		if err != nil {
			response.ResponseError(w, err)
			return
		}

		// 3. Cache miss - check Redis
		tokenKey := auth.GetTokenKey(userId)
		storedToken, err := m.RedisClient.Get(tokenKey)
		if err != nil || storedToken == "" {
			// Token not in Redis or expired
			logx.Errorf("[AuthMiddleware] Token not found in Redis for user %d: %v", userId, err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rawToken := strings.TrimPrefix(token, "Bearer ")
		if storedToken != rawToken {
			// Token mismatch
			logx.Errorf("[AuthMiddleware] Token mismatch for user %d", userId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// 4. Valid token - update local cache
		logx.Infof("[AuthMiddleware] Updating local cache for user %d", userId)

		ctx := context.WithValue(r.Context(), types.CtxKeyUserID, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
