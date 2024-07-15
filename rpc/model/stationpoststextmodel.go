package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StationPostsTextModel = (*customStationPostsTextModel)(nil)

type (
	// StationPostsTextModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStationPostsTextModel.
	StationPostsTextModel interface {
		stationPostsTextModel
		withSession(session sqlx.Session) StationPostsTextModel
	}

	customStationPostsTextModel struct {
		*defaultStationPostsTextModel
	}
)

// NewStationPostsTextModel returns a model for the database table.
func NewStationPostsTextModel(conn sqlx.SqlConn) StationPostsTextModel {
	return &customStationPostsTextModel{
		defaultStationPostsTextModel: newStationPostsTextModel(conn),
	}
}

func (m *customStationPostsTextModel) withSession(session sqlx.Session) StationPostsTextModel {
	return NewStationPostsTextModel(sqlx.NewSqlConnFromSession(session))
}
