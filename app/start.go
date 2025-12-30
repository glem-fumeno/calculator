package app

import (
	"database/sql"

	"github.com/glem-fumeno/calculator/services"
	"github.com/glem-fumeno/calculator/tui"
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
func (s *StartState) GetTitle() string {
	return "Welcome to calculator!"
}
func (s *StartState) GetOptions() []tui.Option {
	return []tui.Option{
		tui.NewOption("C", "Calculate a recipe chain"),
		tui.NewOption("I", "Browse Items"),
		tui.NewOption("R", "Browse Recipes"),
		tui.NewOption("X", "Exit"),
	}
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
