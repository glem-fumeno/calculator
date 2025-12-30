package app

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/services"
)

type StartState struct {
	Services services.Services
}

func NewStartState(database *sql.DB) *StartState {
	return &StartState{
		Services: services.NewServices(database),
	}
}

func (s *StartState) GetError() string {
	return ""
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
		return NewBrowseItemsState(s, s.Services)
	case "R":
		return NewBrowseRecipesState(s, s.Services)
	}
	return nil
}
