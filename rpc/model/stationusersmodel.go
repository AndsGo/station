package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StationUsersModel = (*customStationUsersModel)(nil)

type (
	// StationUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStationUsersModel.
	StationUsersModel interface {
		stationUsersModel
		withSession(session sqlx.Session) StationUsersModel
	}

	customStationUsersModel struct {
		*defaultStationUsersModel
	}
)

// NewStationUsersModel returns a model for the database table.
func NewStationUsersModel(conn sqlx.SqlConn) StationUsersModel {
	return &customStationUsersModel{
		defaultStationUsersModel: newStationUsersModel(conn),
	}
}

func (m *customStationUsersModel) withSession(session sqlx.Session) StationUsersModel {
	return NewStationUsersModel(sqlx.NewSqlConnFromSession(session))
}
