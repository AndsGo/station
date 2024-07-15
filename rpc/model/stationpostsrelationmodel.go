package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StationPostsRelationModel = (*customStationPostsRelationModel)(nil)

type (
	// StationPostsRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStationPostsRelationModel.
	StationPostsRelationModel interface {
		stationPostsRelationModel
		withSession(session sqlx.Session) StationPostsRelationModel
	}

	customStationPostsRelationModel struct {
		*defaultStationPostsRelationModel
	}
)

// NewStationPostsRelationModel returns a model for the database table.
func NewStationPostsRelationModel(conn sqlx.SqlConn) StationPostsRelationModel {
	return &customStationPostsRelationModel{
		defaultStationPostsRelationModel: newStationPostsRelationModel(conn),
	}
}

func (m *customStationPostsRelationModel) withSession(session sqlx.Session) StationPostsRelationModel {
	return NewStationPostsRelationModel(sqlx.NewSqlConnFromSession(session))
}
