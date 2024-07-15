package posts

import (
	"context"
	"rpc/station"

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

	info, err := l.svcCtx.PostsRpc.AddPosts(l.ctx, &station.PostsInfo{
		Id:         req.Id,
		Content:    req.Content,
		Title:      req.Title,
		Source:     req.Source,
		Author:     req.Author,
		ThrownNum:  req.ThrownNum,
		Categories: req.Categories,
		CreateTime: req.CreateTime,
	})
	if err != nil {
		return nil, err
	}
	return &types.PostsInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code:    200,
			Message: "success",
		},
		Data: types.PostsInfo{
			Id:         info.Id,
			Content:    info.Content,
			Title:      info.Title,
			Source:     info.Source,
			Author:     info.Author,
			ThrownNum:  info.ThrownNum,
			Categories: info.Categories,
			CreateTime: info.CreateTime,
		},
	}, nil
}
