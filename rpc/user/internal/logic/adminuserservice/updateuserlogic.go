package adminuserservicelogic

import (
	"context"

	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/rpc/user/internal/svc"
	"bookstore/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UserUpdateReq) (*user.UserInfo, error) {
	if in.Info.Id == 0 {
		return nil, errs.ErrUserNotFound.GRPCStatus().Err()
	}

	err := l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var err error
		adminUserModel := l.svcCtx.AdminUserModel.WithSession(session)
		adminUserRoleModel := l.svcCtx.AdminUserRoleModel.WithSession(session)

		oldUser, err := adminUserModel.FindOne(ctx, in.Info.Id)
		if err != nil {
			return err
		}

		err = adminUserModel.Update(ctx, &model.TAdminUser{
			Id:       in.Info.Id,
			Username: in.Info.UserName,
			Nickname: oldUser.Nickname,
			Avatar:   oldUser.Avatar,
			Email:    in.Info.Email,
			Phone:    oldUser.Phone,
			Status:   int64(in.Info.Status),
			Password: oldUser.Password,
		})

		if err != nil {
			if errs.IsDuplicateKeyError(err) {
				return errs.ErrUsernameAlreadyExist.GRPCStatus().Err()
			}
			return err
		}

		err = adminUserRoleModel.DeleteByUserId(ctx, in.Info.Id)
		if err != nil {
			return err
		}

		for _, v := range in.Ids {
			_, err = adminUserRoleModel.Insert(ctx, &model.TAdminUserRole{
				UserId: in.Info.Id,
				RoleId: v,
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	return &user.UserInfo{}, err
}
