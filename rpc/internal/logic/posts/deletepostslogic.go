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
	err := l.svcCtx.PostsModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &station.BaseDataInfo{}, nil
}
