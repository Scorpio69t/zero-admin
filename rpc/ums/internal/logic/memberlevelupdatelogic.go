package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLevelUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelUpdateLogic {
	return &MemberLevelUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelUpdateLogic) MemberLevelUpdate(in *ums.MemberLevelUpdateReq) (*ums.MemberLevelUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &ums.MemberLevelUpdateResp{}, nil
}
