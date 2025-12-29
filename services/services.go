package services

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/services/items"
)

type Services struct {
	Items items.Services
}

func NewServices(database *sql.DB) Services {
	return Services{items.NewServices(database)}
}
