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
	postsTextFieldNames          = builder.RawFieldNames(&PostsText{})
	postsTextRows                = strings.Join(postsTextFieldNames, ",")
	postsTextRowsExpectAutoSet   = strings.Join(stringx.Remove(postsTextFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	postsTextRowsWithPlaceHolder = strings.Join(stringx.Remove(postsTextFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	postsTextModel interface {
		Insert(ctx context.Context, data *PostsText) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*PostsText, error)
		FindOneByPostsId(ctx context.Context, id uint64) (*PostsText, error)
		Update(ctx context.Context, data *PostsText) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultPostsTextModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PostsText struct {
		Id       uint64         `db:"id"`        // 主键
		PostsId  sql.NullInt64  `db:"posts_id"`  //  posts 主键
		Content  sql.NullString `db:"content"`   // html内容
		CreateAt time.Time      `db:"create_at"` // 创建时间
		UpdateAt time.Time      `db:"update_at"` // 修改时间
	}
)

func newPostsTextModel(conn sqlx.SqlConn) *defaultPostsTextModel {
	return &defaultPostsTextModel{
		conn:  conn,
		table: "`posts_text`",
	}
}

func (m *defaultPostsTextModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPostsTextModel) FindOne(ctx context.Context, id uint64) (*PostsText, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postsTextRows, m.table)
	var resp PostsText
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


func (m *defaultPostsTextModel) FindOneByPostsId(ctx context.Context, id uint64) (*PostsText, error){
	query := fmt.Sprintf("select %s from %s where `posts_id` = ? limit 1", postsTextRows, m.table)
	var resp PostsText
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
func (m *defaultPostsTextModel) Insert(ctx context.Context, data *PostsText) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, postsTextRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PostsId, data.Content)
	return ret, err
}

func (m *defaultPostsTextModel) Update(ctx context.Context, data *PostsText) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, postsTextRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PostsId, data.Content, data.Id)
	return err
}

func (m *defaultPostsTextModel) tableName() string {
	return m.table
}
