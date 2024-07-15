package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StationModel = (*customStationModel)(nil)

type (
	// StationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStationModel.
	StationModel interface {
		stationModel
		withSession(session sqlx.Session) StationModel
	}

	customStationModel struct {
		*defaultStationModel
	}
)

// NewStationModel returns a model for the database table.
func NewStationModel(conn sqlx.SqlConn) StationModel {
	return &customStationModel{
		defaultStationModel: newStationModel(conn),
	}
}

func (m *customStationModel) withSession(session sqlx.Session) StationModel {
	return NewStationModel(sqlx.NewSqlConnFromSession(session))
}
