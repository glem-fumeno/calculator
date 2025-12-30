package app

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/services"
)

type StartState struct {
	*StateData
}

func NewStartState(database *sql.DB) *StartState {
	return &StartState{
		&StateData{Services: services.NewServices(database)},
	}
}

func (s *StartState) GetOptions() Options {
	return NewOptions(
		NewLine("Welcome to calculator!"),
		NewOption("C", "Calculate a recipe chain"),
		NewOption("I", "Browse Items"),
		NewOption("R", "Browse Recipes"),
		NewOption("X", "Exit"),
	)
}
func (s *StartState) Run(option string) State {
	switch option {
	case "I":
		return NewBrowseItemsState(s)
	case "R":
		return NewBrowseRecipesState(s)
	}
	return nil
}
