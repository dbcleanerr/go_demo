package db

import "database/sql"

type Story struct {
	queries *Queries
	db      *sql.DB
}

func NewStore(db *sql.DB) *Story {
	return &Story{
		queries: New(db),
		db:      db,
	}
}
