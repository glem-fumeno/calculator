package queries

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/queries/items"
	_ "github.com/mattn/go-sqlite3"
)

type Queries struct {
	database   *sql.DB
	connection *sql.Tx

	Items items.Queries
}

func NewQueries(path string) (*Queries, error) {
	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	connection, err := database.Begin()
	if err != nil {
		database.Close()
		return nil, err
	}
	return &Queries{database, connection, items.NewQueries(connection)}, nil
}

func (q *Queries) Close(err *error) error {
	defer q.database.Close()
	if *err != nil {
		return q.connection.Rollback()
	}
	return q.connection.Commit()
}
