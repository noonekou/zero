package logic

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/rpc/add/adder"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	rsp, err := l.svcCtx.Adder.Add(l.ctx, &adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	})

	logx.Infof("AddLogic.Add: rsp: %v, err: %v", rsp, err)

	if err != nil {
		return
	}

	return &types.AddResp{
		Ok: rsp.Ok,
	}, nil
}
