package posts

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增Posts
func NewAddPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostsLogic {
	return &AddPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostsLogic) AddPosts(req *types.PostsInfo) (resp *types.PostsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
