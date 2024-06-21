package posts

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除Posts
func NewDeletePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostsLogic {
	return &DeletePostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostsLogic) DeletePosts(req *types.IDPathReq) (resp *types.BaseDataInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
