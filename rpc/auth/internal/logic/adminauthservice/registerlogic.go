package adminauthservicelogic

import (
	"context"
	"database/sql"

	common "bookstore/common/auth"
	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *auth.RegisterReq) (*auth.RegisterResp, error) {
	if (in.Username == "" || in.Password == "") && in.ConfirmPassword == "" {
		return nil, errs.ErrUsernameOrPasswordIsEmpty.GRPCStatus().Err()
	}

	if in.RoleId == 0 {
		return nil, errs.ErrRoleNotFound.GRPCStatus().Err()
	}

	var token string

	err := l.svcCtx.Conn.TransactCtx(l.ctx, func(txCtx context.Context, conn sqlx.Session) error {

		user, err := l.svcCtx.AdminUserModel.FindOneByUsername(l.ctx, in.Username)
		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if user != nil {
			return errs.ErrUsernameAlreadyExist.GRPCStatus().Err()
		}

		_, err = l.svcCtx.RoleModel.FindOne(l.ctx, in.RoleId)
		if err != nil {
			return errs.ErrRoleNotFound.GRPCStatus().Err()
		}

		_, err = l.svcCtx.AdminUserModel.Insert(l.ctx, &model.TAdminUser{
			Username: in.Username,
			Password: in.Password,
			Nickname: "",
			Avatar:   "",
			Email:    "",
			Phone:    "",
			Status:   1,
		})

		if err != nil {
			return err
		}

		user, err = l.svcCtx.AdminUserModel.FindOneByUsername(l.ctx, in.Username)
		if err != nil {
			return err
		}

		_, err = l.svcCtx.AdminUserRoleModel.Insert(l.ctx, &model.TAdminUserRole{
			UserId: user.Id,
			RoleId: in.RoleId,
		})

		if err != nil {
			return err
		}

		token, err = common.GenerateToken(l.svcCtx.Config.Authorization.AccessSecret, l.svcCtx.Config.Authorization.AccessExpire, user.Id)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &auth.RegisterResp{
		Token: token,
	}, nil
}
