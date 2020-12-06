package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryAddLogic {
	return ProductCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryAddLogic) ProductCategoryAdd(req types.AddProductCategoryReq) (*types.AddProductCategoryResp, error) {
	_, err := l.svcCtx.Pms.ProductCategoryAdd(l.ctx, &pmsclient.ProductCategoryAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddProductCategoryResp{}, nil
}
