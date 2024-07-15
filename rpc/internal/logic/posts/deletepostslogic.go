package postslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostsLogic {
	return &DeletePostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostsLogic) DeletePosts(in *station.IDPathReq) (*station.BaseDataInfo, error) {
	// todo: add your logic here and delete this line

	return &station.BaseDataInfo{}, nil
}
