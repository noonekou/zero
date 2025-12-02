// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"context"

	"bookstore/admin/internal/svc"
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
	// todo: add your logic here and delete this line

	return nil
}
