package app

import (
	"github.com/glem-fumeno/calculator/schemas"
)

type AddRecipeState struct {
	*StateData

	recipe schemas.DBRecipe
}

func NewAddRecipeState(parent State, name string) *AddRecipeState {
	return &AddRecipeState{
		NewStateData(parent),
		schemas.DBRecipe{RecipeName: name},
	}
}

func (s *AddRecipeState) GetOptions() Options {
	return NewOptions(
		NewLine("Adding an recipe"),
		NewOption("N", "Name: %s", s.recipe.RecipeName),
		NewOption("S", "Save and go back"),
		NewOption("B", "Back"),
	)
}
func (s *AddRecipeState) Run(option string) State {
	switch option {
	case "N":
		s.recipe.RecipeName = GetInput("Name")
	case "S":
		err := s.Services.Recipes.Create(s.recipe)
		if err == nil {
			return s.Parent
		}
		s.error = err.Error()
	case "B":
		return s.Parent
	}
	return s
}
