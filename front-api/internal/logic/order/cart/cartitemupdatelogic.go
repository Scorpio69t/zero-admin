package logic

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartItemUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartItemUpdateLogic {
	return CartItemUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartItemUpdateLogic) CartItemUpdate(req types.UpdateCartItemReq) (*types.UpdateCartItemResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateCartItemResp{}, nil
}
