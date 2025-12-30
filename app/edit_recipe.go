package app

import (
	"fmt"

	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
	"github.com/glem-fumeno/calculator/tui"
)

type EditRecipeState struct {
	Parent   State
	Services services.Services

	recipeName string
	recipe     schemas.DBRecipe
	error    string
}

func NewEditRecipeState(
	parent State,
	s services.Services,
	recipe schemas.DBRecipe,
) *EditRecipeState {
	return &EditRecipeState{
		Parent:   parent,
		Services: s,
		recipeName: recipe.RecipeName,
		recipe:     recipe,
	}
}

func (s *EditRecipeState) GetError() string {
	return s.error
}
func (s *EditRecipeState) GetTitle() string {
	return fmt.Sprintf("Editing %s", s.recipeName)
}
func (s *EditRecipeState) GetOptions() []tui.Option {
	return []tui.Option{
		tui.NewOption("N", "Name: %s", s.recipe.RecipeName),
		tui.NewOption("D", "Delete"),
		tui.NewOption("S", "Save and go back"),
		tui.NewOption("B", "Back"),
	}
}
func (s *EditRecipeState) Run(option string) State {
	switch option {
	case "N":
		s.recipe.RecipeName = tui.GetInput("Name")
	case "D":
		err := s.Services.Recipes.Delete(s.recipeName)
		if err == nil {
			return s.Parent
		}
		s.error = err.Error()
	case "S":
		err := s.Services.Recipes.Update(s.recipeName, s.recipe)
		if err == nil {
			return s.Parent
		}
		s.error = err.Error()
	case "B":
		return s.Parent
	}
	return s
}
