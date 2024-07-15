package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DeliveryLogModel = (*customDeliveryLogModel)(nil)

type (
	// DeliveryLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeliveryLogModel.
	DeliveryLogModel interface {
		deliveryLogModel
		withSession(session sqlx.Session) DeliveryLogModel
	}

	customDeliveryLogModel struct {
		*defaultDeliveryLogModel
	}
)

// NewDeliveryLogModel returns a model for the database table.
func NewDeliveryLogModel(conn sqlx.SqlConn) DeliveryLogModel {
	return &customDeliveryLogModel{
		defaultDeliveryLogModel: newDeliveryLogModel(conn),
	}
}

func (m *customDeliveryLogModel) withSession(session sqlx.Session) DeliveryLogModel {
	return NewDeliveryLogModel(sqlx.NewSqlConnFromSession(session))
}
