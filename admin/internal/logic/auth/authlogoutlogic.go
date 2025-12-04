// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"bookstore/admin/internal/svc"
	"bookstore/admin/internal/types"
	common "bookstore/common/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 退出登录
func NewAuthLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogoutLogic {
	return &AuthLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogoutLogic) AuthLogout() error {
	// Get userId from context (set by auth middleware if present)
	// If no middleware, return success (already logged out or no valid token)
	userIdVal := l.ctx.Value(types.CtxKeyUserID)
	if userIdVal == nil {
		// No valid token in context, consider already logged out
		logx.Error("No valid token in context")
		return nil
	}

	userId, ok := userIdVal.(int64)
	if !ok {
		logx.Error("[AuthLogout] Failed to cast userId to int64")
		return nil
	}

	// Delete token from Redis
	tokenKey := common.GetTokenKey(userId)
	_, err := l.svcCtx.RedisClient.Del(tokenKey)
	if err != nil {
		logx.Errorf("[AuthLogout] Failed to delete token from Redis: %v", err)
		// Continue anyway - logout should succeed even if Redis fails
	} else {
		logx.Infof("[AuthLogout] Deleted token from Redis for user %d", userId)
	}

	return nil
}
