package main

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/services"
	_ "github.com/mattn/go-sqlite3"
)

var (
	service services.Services
)

func main() {
	database, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	service = services.NewServices(database)
	for {
		option := getOption(
			"Welcome to calculator!",
			"",
			[]Option{
				{"C", "Calculate a recipe chain"},
				{"I", "Browse Items"},
				{"R", "Browse Recipes"},
				{"X", "Exit"},
			},
		)
		switch option {
		case "C":
			calculateRecipeChain()
		case "I":
			browseItems()
		case "R":
			calculateRecipeChain()
		case "X":
			return
		}
	}
}
