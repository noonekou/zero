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
			Password: "827ccb0eea8a706c4c34a16891f84e7b", // 12345
		})

		if err != nil {
			if errs.IsDuplicateKeyError(err) {
				return errs.ErrUsernameAlreadyExist.GRPCStatus().Err()
			}
			return err
		}

		adminUserRoleModel := l.svcCtx.AdminUserRoleModel.WithSession(session)

		// 批量插入用户角色关系
		if len(in.Ids) > 0 {
			userRoles := make([]*model.TAdminUserRole, 0, len(in.Ids))
			for _, roleId := range in.Ids {
				userRoles = append(userRoles, &model.TAdminUserRole{
					UserId: uid,
					RoleId: roleId,
				})
			}

			err = adminUserRoleModel.BatchInsert(ctx, userRoles)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return &user.UserInfo{}, err
}
