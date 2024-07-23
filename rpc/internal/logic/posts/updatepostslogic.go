package postslogic

import (
	"context"
	"database/sql"

	"rpc/internal/svc"
	"rpc/model"
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
	info := &model.Posts{
		Id:         in.Id,
		Title:      sql.NullString{String: in.Title, Valid: true},
		Source:     sql.NullString{String: in.Source, Valid: true},
		Author:     sql.NullInt64{Int64: in.Author, Valid: true},
		ThrownNum:  sql.NullInt64{Int64: in.ThrownNum, Valid: true},
		Categories: sql.NullString{String: in.Categories, Valid: true},
	}
	text := &model.PostsText{
		PostsId: sql.NullInt64{Int64: int64(in.Id), Valid: true},
		Content: sql.NullString{String: in.Content, Valid: true},
	}
	err := l.svcCtx.PostsModel.UpdatePosts(l.ctx, info, text)
	if err != nil {
		return nil, err
	}
	result := &station.PostsInfo{
		Id:         in.Id,
		Title:      info.Title.String,
		Source:     info.Source.String,
		Author:     info.Author.Int64,
		ThrownNum:  info.ThrownNum.Int64,
		Categories: info.Categories.String,
		Content:    text.Content.String,
	}
	return result, nil
}
