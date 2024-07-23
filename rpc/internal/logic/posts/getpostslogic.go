package postslogic

import (
	"context"

	"rpc/internal/svc"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文章帖子详情
func (l *GetPostsLogic) GetPosts(in *station.IDPathReq) (*station.PostsInfo, error) {
	// 1.获取基础信息
	res, err := l.svcCtx.PostsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	// 2.获取详情
	text, err := l.svcCtx.PostsTextModel.FindOneByPostsId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	// 3.封装
	resp := &station.PostsInfo{
		Id:         res.Id,
		Title:      res.Title.String,
		Source:     res.Source.String,
		Author:     res.Author.Int64,
		CreateTime: res.CreateAt.Unix(),
		ThrownNum:  res.ThrownNum.Int64,
		Categories: res.Categories.String,
		Content:    text.Content.String,
	}
	return resp, nil
}
