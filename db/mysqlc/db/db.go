package db

import (
	"context"
	"database/sql"
)

type DBTX interface {
	// ExecContext 不返回记录
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	// QueryContext 返回记录集
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	// QueryRowContext 返回单条记录
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

// Queries 所有DB的操作
type Queries struct {
	db DBTX
}

// New 初始化 Queries, 所有实现了 DBTX 接口的方法都可以
func New(db DBTX) *Queries {
	return &Queries{db: db}
}

// WithTx db事务使用的方法, tx 也实现了 DBTX 接口
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{db: tx}
}
