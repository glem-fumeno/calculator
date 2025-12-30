package app

import (
	"fmt"

	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
)

type EditRecipeState struct {
	Parent   State
	Services services.Services

	recipeName string
	recipe     schemas.DBRecipe
	error      string
}

func NewEditRecipeState(
	parent State,
	s services.Services,
	recipe schemas.DBRecipe,
) *EditRecipeState {
	return &EditRecipeState{
		Parent:     parent,
		Services:   s,
		recipeName: recipe.RecipeName,
		recipe:     recipe,
	}
}

func (s *EditRecipeState) GetError() string {
	return s.error
}
func (s *EditRecipeState) GetOptions() Options {
	ingredients, products, err := s.Services.Recipes.ReadItems(s.recipeName)
	if err != nil {
		panic(err)
	}
	options := NewOptions(
		NewLine("Editing %s", s.recipeName),
		NewOption("N", "Name: %s", s.recipe.RecipeName),
		NewOption("S", "Save name"),
		NewLine("Ingredients"),
		NewOption("IA", "Add ingredient"),
	)
	for i, item := range ingredients {
		options = options.Add(
			NewOption(fmt.Sprintf("I%d", i+1), "%s (%d)", item.ItemName, item.Quantity),
		)
	}
	options = options.Add(
		NewLine("Products"),
		NewOption("PA", "Add Product"),
	)
	for i, item := range products {
		options = options.Add(
			NewOption(fmt.Sprintf("P%d", i+1), "%s (%d)", item.ItemName, item.Quantity),
		)
	}
	return options.Add(
		NewOption("D", "Delete"),
		NewOption("B", "Back"),
	)
}
func (s *EditRecipeState) Run(option string) State {
	switch option {
	case "N":
		s.recipe.RecipeName = GetInput("Name")
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
