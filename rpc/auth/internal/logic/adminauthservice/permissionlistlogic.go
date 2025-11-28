package adminauthservicelogic

import (
	"context"

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
	resources, err := l.svcCtx.ResourceModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	permissions, err := l.svcCtx.PermissionModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	list := make([]*auth.Permission, 0)

	rootResource := lo.Filter(resources, func(v model.TResource, index int) bool {
		return v.ParentCode == 0
	})

	for _, v := range rootResource {
		subResource := lo.Filter(resources, func(m model.TResource, index int) bool {
			return m.ParentCode == v.Code
		})
		children := make([]*auth.Permission, 0)
		for _, s := range subResource {
			curPermission := lo.Filter(permissions, func(m model.TPermission, index int) bool {
				return m.ResourceName == s.Name
			})

			for _, p := range curPermission {
				children = append(children, &auth.Permission{Id: p.Id, Code: int32(s.Code), Description: p.Description.String, ParentCode: int32(s.ParentCode), Children: nil, CreatedAt: s.CreatedAt.Unix(), UpdatedAt: s.UpdatedAt.Unix()})
			}

		}
		pId := v.Id
		if len(children) > 0 {
			pId = 0
		}
		permission := auth.Permission{Id: pId, Code: int32(v.Code), Description: v.Description, ParentCode: int32(v.ParentCode), Children: children, CreatedAt: v.CreatedAt.Unix(), UpdatedAt: v.UpdatedAt.Unix()}
		list = append(list, &permission)
	}

	return &auth.PermissionListResp{
		List: list,
	}, nil
}
