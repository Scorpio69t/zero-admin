package logic

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogic {
	return MemberLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogic) MemberLogin(req types.MemberLoginReq) (*types.MemberLoginResp, error) {
	// todo: add your logic here and delete this line

	return &types.MemberLoginResp{}, nil
}
