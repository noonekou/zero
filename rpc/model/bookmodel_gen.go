// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.4

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	bookFieldNames          = builder.RawFieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "`book`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	bookModel interface {
		Insert(ctx context.Context, data *Book) (sql.Result, error)
		FindOne(ctx context.Context, book string) (*Book, error)
		Update(ctx context.Context, data *Book) error
		Delete(ctx context.Context, book string) error
	}

	defaultBookModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Book struct {
		Book  string `db:"book"`  // book name
		Price int64  `db:"price"` // book price
	}
)

func newBookModel(conn sqlx.SqlConn) *defaultBookModel {
	return &defaultBookModel{
		conn:  conn,
		table: "`book`",
	}
}

func (m *defaultBookModel) Delete(ctx context.Context, book string) error {
	query := fmt.Sprintf("delete from %s where `book` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, book)
	return err
}

func (m *defaultBookModel) FindOne(ctx context.Context, book string) (*Book, error) {
	query := fmt.Sprintf("select %s from %s where `book` = ? limit 1", bookRows, m.table)
	var resp Book
	err := m.conn.QueryRowCtx(ctx, &resp, query, book)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) Insert(ctx context.Context, data *Book) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bookRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Book, data.Price)
	return ret, err
}

func (m *defaultBookModel) Update(ctx context.Context, data *Book) error {
	query := fmt.Sprintf("update %s set %s where `book` = ?", m.table, bookRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Price, data.Book)
	return err
}

func (m *defaultBookModel) tableName() string {
	return m.table
}
