// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	stationFieldNames          = builder.RawFieldNames(&Station{})
	stationRows                = strings.Join(stationFieldNames, ",")
	stationRowsExpectAutoSet   = strings.Join(stringx.Remove(stationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	stationRowsWithPlaceHolder = strings.Join(stringx.Remove(stationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	stationModel interface {
		Insert(ctx context.Context, data *Station) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Station, error)
		Update(ctx context.Context, data *Station) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultStationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Station struct {
		Id           uint64          `db:"id"`            // 主键
		DomainName   sql.NullString  `db:"domain_name"`   // 域名
		Ip           sql.NullString  `db:"ip"`            // ip
		DomainYear   sql.NullInt64   `db:"domain_year"`   // 域名年份
		GoogleWeight sql.NullFloat64 `db:"google_weight"` // 谷歌权重
		Type         sql.NullString  `db:"type"`          // 网站类型
		Industry     sql.NullString  `db:"industry"`      // 网站行业
		ArticlesNum  sql.NullInt64   `db:"articles_num"`  // 文章数量
		UserName     sql.NullString  `db:"user_name"`     // 账号名
		PassWord     sql.NullString  `db:"pass_word"`     // 密码
		CreateAt     time.Time       `db:"create_at"`     // 创建时间
		UpdateAt     time.Time       `db:"update_at"`     // 修改时间
	}
)

func newStationModel(conn sqlx.SqlConn) *defaultStationModel {
	return &defaultStationModel{
		conn:  conn,
		table: "`station`",
	}
}

func (m *defaultStationModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultStationModel) FindOne(ctx context.Context, id uint64) (*Station, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", stationRows, m.table)
	var resp Station
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStationModel) Insert(ctx context.Context, data *Station) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, stationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DomainName, data.Ip, data.DomainYear, data.GoogleWeight, data.Type, data.Industry, data.ArticlesNum, data.UserName, data.PassWord)
	return ret, err
}

func (m *defaultStationModel) Update(ctx context.Context, data *Station) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, stationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DomainName, data.Ip, data.DomainYear, data.GoogleWeight, data.Type, data.Industry, data.ArticlesNum, data.UserName, data.PassWord, data.Id)
	return err
}

func (m *defaultStationModel) tableName() string {
	return m.table
}