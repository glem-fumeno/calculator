package items

import "database/sql"

type Services struct {
	database *sql.DB
}

func NewServices(database *sql.DB) Services {
	return Services{database}
}
