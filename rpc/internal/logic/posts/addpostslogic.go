package postslogic

import (
	"context"
	"database/sql"

	"rpc/internal/svc"
	"rpc/model"
	"rpc/station"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostsLogic {
	return &AddPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPostsLogic) AddPosts(in *station.PostsInfo) (*station.PostsInfo, error) {
	// todo: add your logic here and delete this line
	info := &model.Posts{
		Title:      sql.NullString{String: in.Title, Valid: true},
		Source:     sql.NullString{String: in.Source, Valid: true},
		Author:     sql.NullInt64{Int64: in.Author, Valid: true},
		ThrownNum:  sql.NullInt64{Int64: in.ThrownNum, Valid: true},
		Categories: sql.NullString{String: in.Categories, Valid: true},
	}
	r, err := l.svcCtx.PostsModel.Insert(l.ctx, info)
	if err != nil {
		return nil, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}
	result := &station.PostsInfo{
		Id:         uint64(id),
		Title:      info.Title.String,
		Source:     info.Source.String,
		Author:     info.Author.Int64,
		ThrownNum:  info.ThrownNum.Int64,
		Categories: info.Categories.String,
	}
	return result, nil
}
