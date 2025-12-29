package queries

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/queries/items"
)

type Queries struct {
	connection *sql.Tx

	Items items.Queries
}

func NewQueries(database *sql.DB) (*Queries, error) {
	connection, err := database.Begin()
	if err != nil {
		return nil, err
	}
	return &Queries{connection, items.NewQueries(connection)}, nil
}

func (q *Queries) Rollback() error {
	return q.connection.Rollback()
}

func (q *Queries) Commit() error {
	return q.connection.Commit()
}
