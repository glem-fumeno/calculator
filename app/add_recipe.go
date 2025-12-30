package app

import (
	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
)

type AddRecipeState struct {
	Parent   State
	Services services.Services

	recipe schemas.DBRecipe
	error  string
}

func NewAddRecipeState(
	parent State,
	s services.Services,
	name string,
) *AddRecipeState {
	return &AddRecipeState{
		Parent:   parent,
		Services: s,
		recipe:   schemas.DBRecipe{RecipeName: name},
	}
}

func (s *AddRecipeState) GetError() string {
	return s.error
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
