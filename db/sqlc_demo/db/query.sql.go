// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createAuthor = `-- name: CreateAuthor :exec
INSERT INTO authors (name, bio) VALUES ($1, $2)
`

type CreateAuthorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) error {
	_, err := q.exec(ctx, q.createAuthorStmt, createAuthor, arg.Name, arg.Bio)
	return err
}

const createGetAuthor = `-- name: CreateGetAuthor :one
INSERT INTO authors (name, bio) VALUES ($1, $2)
RETURNING id, name, bio
`

type CreateGetAuthorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateGetAuthor(ctx context.Context, arg CreateGetAuthorParams) (Author, error) {
	row := q.queryRow(ctx, q.createGetAuthorStmt, createGetAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
 WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteAuthorStmt, deleteAuthor, id)
	return err
}

const deleteGetAuthor = `-- name: DeleteGetAuthor :one
DELETE FROM authors
 WHERE id = $1
RETURNING id, name, bio
`

func (q *Queries) DeleteGetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.queryRow(ctx, q.deleteGetAuthorStmt, deleteGetAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio
  FROM authors
 WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.queryRow(ctx, q.getAuthorStmt, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listAuthorByIDs = `-- name: ListAuthorByIDs :many
SELECT id, name, bio
  FROM authors
 WHERE id = ANY($1::int[])
`

func (q *Queries) ListAuthorByIDs(ctx context.Context, dollar_1 []int32) ([]Author, error) {
	rows, err := q.query(ctx, q.listAuthorByIDsStmt, listAuthorByIDs, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio FROM authors
 ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.query(ctx, q.listAuthorsStmt, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE authors
   SET name = $2,
       bio = $3
 WHERE id = $1
`

type UpdateAuthorParams struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.exec(ctx, q.updateAuthorStmt, updateAuthor, arg.ID, arg.Name, arg.Bio)
	return err
}

const updateGetAuthor = `-- name: UpdateGetAuthor :one
UPDATE authors
   SET name = $2,
       bio = $3
 WHERE id = $1
RETURNING id, name, bio
`

type UpdateGetAuthorParams struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

func (q *Queries) UpdateGetAuthor(ctx context.Context, arg UpdateGetAuthorParams) (Author, error) {
	row := q.queryRow(ctx, q.updateGetAuthorStmt, updateGetAuthor, arg.ID, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}
