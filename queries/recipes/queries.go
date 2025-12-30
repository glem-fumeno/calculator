package recipes

import "database/sql"

type Queries struct {
	connection *sql.Tx
}

func NewQueries(connection *sql.Tx) Queries {
	return Queries{connection}
}
