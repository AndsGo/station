package postslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostsLogic {
	return &UpdatePostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostsLogic) UpdatePosts(in *station.PostsInfo) (*station.PostsInfo, error) {
	// todo: add your logic here and delete this line

	return &station.PostsInfo{}, nil
}
