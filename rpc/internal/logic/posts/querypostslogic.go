package postslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostsLogic {
	return &QueryPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryPostsLogic) QueryPosts(in *station.PostsReq) (*station.PostsListInfo, error) {
	// todo: add your logic here and delete this line

	return &station.PostsListInfo{}, nil
}
