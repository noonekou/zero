package adminauthservicelogic

import (
	"context"
	"database/sql"

	common "bookstore/common/auth"
	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/common/utils"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *auth.LoginReq) (*auth.LoginResp, error) {
	if in.Username == "" || in.Password == "" {
		return nil, errs.ErrUsernameOrPasswordIsEmpty.GRPCStatus().Err()
	}

	logx.Infof("username: %s, password: %s", in.Username, in.Password)
	var tUser *model.TAdminUser
	var err error
	if utils.IsEmail(in.Username) {
		tUser, err = l.svcCtx.AdminUserModel.FindOneByEmailAndPassword(l.ctx, in.Username, in.Password)
	} else {
		tUser, err = l.svcCtx.AdminUserModel.FindOneByUsernameAndPassword(l.ctx, in.Username, in.Password)
	}

	logx.Errorf("tUser: %v", err)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if tUser == nil {
		return nil, errs.ErrUsernameNotExist.GRPCStatus().Err()
	}

	// Check user status (already filtered in query, but double-check for safety)
	if tUser.Status != 1 {
		return nil, errs.ErrUserDisabled.GRPCStatus().Err()
	}

	token, err := common.GenerateToken(l.svcCtx.Config.Authorization.AccessSecret, l.svcCtx.Config.Authorization.AccessExpire, tUser.Id)
	if err != nil {
		return nil, err
	}

	// Store token in Redis
	tokenKey := common.GetTokenKey(tUser.Id)
	err = l.svcCtx.RedisClient.Setex(tokenKey, token, int(l.svcCtx.Config.Authorization.AccessExpire))
	if err != nil {
		logx.Errorf("Failed to store token in Redis: %v", err)
		// Continue even if Redis fails (fallback to JWT validation)
	}

	return &auth.LoginResp{Token: token}, nil
}
