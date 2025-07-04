package logic

import (
	"context"

	"bookstore/rpc/check/internal/svc"
	"bookstore/rpc/check/pb/check"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	resp, err := l.svcCtx.Model.FindOne(l.ctx, in.Book)

	if err != nil {
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
}
