package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PostsModel = (*customPostsModel)(nil)

type (
	// PostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsModel.
	PostsModel interface {
		postsModel
		withSession(session sqlx.Session) PostsModel
	}

	customPostsModel struct {
		*defaultPostsModel
	}
)

// NewPostsModel returns a model for the database table.
func NewPostsModel(conn sqlx.SqlConn) PostsModel {
	return &customPostsModel{
		defaultPostsModel: newPostsModel(conn),
	}
}

func (m *customPostsModel) withSession(session sqlx.Session) PostsModel {
	return NewPostsModel(sqlx.NewSqlConnFromSession(session))
}
