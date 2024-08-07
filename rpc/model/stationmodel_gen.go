// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"rpc/client/station"
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
		FindAll(ctx context.Context) ([]*Station, error)
		QueryStation(ctx context.Context, data *station.StationReq) ([]*Station,uint64, error)
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
		Status       sql.NullInt64   `db:"status"`        // 状态 0 正常 1 失效 2 删除
		Creater      sql.NullString  `db:"creater"`       // 创建人
		CreateAt     time.Time       `db:"create_at"`     // 创建时间
		Modifier     sql.NullString  `db:"modifier"`      // 修改人
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
	// 这里进行软删除,将状态修改为2
	query := fmt.Sprintf("update %s set `status` = 2 where `id` = ?", m.table)
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
// 获取所有正常的站点
func (m *defaultStationModel) FindAll(ctx context.Context) ([]*Station, error){
	resp := make([]*Station, 0)
	err := m.conn.QueryRows(&resp, fmt.Sprintf("select %s from %s where status = 0", stationRows, m.table))
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 分页查询
func (m *defaultStationModel) QueryStation(ctx context.Context, data *station.StationReq) ([]*Station,uint64, error){
	// 1.构建查询条件
	query := "where 1=1"
	params :=make([]interface{},0,0)
	if data.DomainName!=""{ 
		query += " and domain_name like ?"
		params = append(params,  "%"+data.DomainName+"%")
	}
	if data.Ip!=""{ 
		query += " and ip like ?"
		params = append(params,  "%"+data.Ip+"%")
	}
	if data.DomainYear!=0{ 
		query += " and domain_year = ?"
		params = append(params,  data.DomainYear)
	}
	if data.GoogleWeight!=""{ 
		query += " and google_weight = ?"
		params = append(params,  data.GoogleWeight)
	}
	if data.Type!=""{ 
		query += " and type = ?"
		params = append(params,  data.Type)
	}
	if data.Industry!=""{ 
		query += " and industry = ?"
		params = append(params,  data.Industry)
	}
	query += " and status = ?"
		params = append(params,  data.Status)
	// 2.查数量
	total := 0
	items := make([]*Station, 0)
	err := m.conn.QueryRow(&total,fmt.Sprintf("select count(1) from %s ", m.table) + query, params...)
	if err!=nil {
		return nil,0, err
	}
	// 没有记录
	if total == 0 {
		return items, uint64(total), nil
	}
	//分页
	query += "  order by id desc LIMIT ? OFFSET ?"
	params = append(params, data.PageInfo.PageSize)
	params = append(params, data.PageInfo.PageSize*(data.PageInfo.Page-1))
	// 3.分页查询数据
	err = m.conn.QueryRows(&items, fmt.Sprintf("select %s from %s", stationRows, m.table) + query, params...)
	if err!=nil {
		return nil,0, err
	}
	return items,uint64(total),nil
}

func (m *defaultStationModel) Insert(ctx context.Context, data *Station) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, stationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DomainName, data.Ip, data.DomainYear, data.GoogleWeight, data.Type, data.Industry, data.ArticlesNum, data.UserName, data.PassWord, data.Status, data.Creater, data.Modifier)
	return ret, err
}

func (m *defaultStationModel) Update(ctx context.Context, data *Station) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, stationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DomainName, data.Ip, data.DomainYear, data.GoogleWeight, data.Type, data.Industry, data.ArticlesNum, data.UserName, data.PassWord, data.Status, data.Creater, data.Modifier, data.Id)
	return err
}

func (m *defaultStationModel) tableName() string {
	return m.table
}
