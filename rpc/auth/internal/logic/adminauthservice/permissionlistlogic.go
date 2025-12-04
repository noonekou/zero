package adminauthservicelogic

import (
	"cmp"
	"context"
	"slices"

	"bookstore/common/model"
	"bookstore/rpc/auth/auth"
	"bookstore/rpc/auth/internal/svc"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionListLogic {
	return &PermissionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionListLogic) PermissionList(in *auth.PermissionListReq) (*auth.PermissionListResp, error) {

	all, err := l.svcCtx.PermissionModel.FindParentAll(l.ctx)
	if err != nil {
		return nil, err
	}

	group := lo.GroupBy(all, func(v model.TPermissionParentData) int {
		return v.ParentCode
	})

	var iter func(pCode int) []*auth.Permission
	iter = func(pCode int) []*auth.Permission {
		if _, ok := group[pCode]; !ok {
			return []*auth.Permission{}
		}
		children := make([]*auth.Permission, 0)
		for _, v := range group[pCode] {
			children = append(children, &auth.Permission{
				Id:          v.Id,
				Code:        int32(v.Code),
				Description: v.Description,
				ParentCode:  int32(v.ParentCode),
				Children:    iter(int(v.Code)),
				CreatedAt:   v.CreatedAt.Unix(),
				UpdatedAt:   v.UpdatedAt.Unix(),
			})
		}

		slices.SortFunc(children, func(a, b *auth.Permission) int {
			return cmp.Compare(a.Id, b.Id)
		})

		return children
	}

	temp := make([]*auth.Permission, 0)
	for _, v := range group[0] {
		temp = append(temp, &auth.Permission{
			Id:          v.Id,
			Code:        int32(v.Code),
			Description: v.PDescription,
			ParentCode:  int32(v.ParentCode),
			Children:    iter(int(v.Code)),
			CreatedAt:   v.CreatedAt.Unix(),
			UpdatedAt:   v.UpdatedAt.Unix(),
		})
	}

	return &auth.PermissionListResp{
		List: temp,
	}, nil
}
