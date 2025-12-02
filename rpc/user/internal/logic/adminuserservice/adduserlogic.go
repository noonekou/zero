package adminuserservicelogic

import (
	"context"
	"errors"

	errs "bookstore/common/error"
	"bookstore/common/model"
	"bookstore/rpc/user/internal/svc"
	"bookstore/rpc/user/user"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.UserUpdateReq) (*user.UserInfo, error) {
	err := l.svcCtx.Conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var err error
		adminUserModel := l.svcCtx.AdminUserModel.WithSession(session)

		uid, err := adminUserModel.InsertWithId(ctx, &model.TAdminUser{
			Username: in.Info.UserName,
			Nickname: in.Info.NickName,
			Avatar:   in.Info.Avatar,
			Email:    in.Info.Email,
			Phone:    in.Info.Phone,
			Status:   int64(in.Info.Status),
		})

		if err != nil {
			var pgErr *pq.Error
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return errs.ErrUsernameAlreadyExist.GRPCStatus().Err()
			}
			return err
		}

		adminUserRoleModel := l.svcCtx.AdminUserRoleModel.WithSession(session)
		for _, v := range in.Ids {
			_, err = adminUserRoleModel.Insert(ctx, &model.TAdminUserRole{
				UserId: uid,
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
