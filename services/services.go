package services

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/services/items"
	"github.com/glem-fumeno/calculator/services/recipes"
)

type Services struct {
	Items items.Services
	Recipes recipes.Services
}

func NewServices(database *sql.DB) Services {
	return Services{
		items.NewServices(database),
		recipes.NewServices(database),
	}
}
