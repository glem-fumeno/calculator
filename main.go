package main

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	a := app.App{State: app.NewStartState(database)}
	a.Run()
}
