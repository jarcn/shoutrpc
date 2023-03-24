package logic

import (
	"context"

	"github.com/jarcn/shoutrpc/internal/svc"
	"github.com/jarcn/shoutrpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	// 手动代码开始
	res, err := l.svcCtx.Model.FindOne(l.ctx, in.Shorten)
	if err != nil {
		return nil, err
	}

	return &transform.ExpandResp{
		Url: res.Url,
	}, nil
	// 手动代码结束
}
