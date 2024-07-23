package posts

import (
	"context"
	"rpc/station"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改Posts
func NewUpdatePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostsLogic {
	return &UpdatePostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostsLogic) UpdatePosts(req *types.PostsInfo) (resp *types.PostsInfoResp, err error) {
	info, err := l.svcCtx.PostsRpc.UpdatePosts(l.ctx, &station.PostsInfo{
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
			Message: types.Success,
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
