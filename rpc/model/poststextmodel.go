package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PostsTextModel = (*customPostsTextModel)(nil)

type (
	// PostsTextModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsTextModel.
	PostsTextModel interface {
		postsTextModel
		withSession(session sqlx.Session) PostsTextModel
	}

	customPostsTextModel struct {
		*defaultPostsTextModel
	}
)

// NewPostsTextModel returns a model for the database table.
func NewPostsTextModel(conn sqlx.SqlConn) PostsTextModel {
	return &customPostsTextModel{
		defaultPostsTextModel: newPostsTextModel(conn),
	}
}

func (m *customPostsTextModel) withSession(session sqlx.Session) PostsTextModel {
	return NewPostsTextModel(sqlx.NewSqlConnFromSession(session))
}
