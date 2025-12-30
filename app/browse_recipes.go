package app

import (
	"fmt"
	"strconv"

	"github.com/glem-fumeno/calculator/schemas"
)

type BrowseRecipesState struct {
	*StateData
	recipes []schemas.DBRecipe
}

func NewBrowseRecipesState(parent State) *BrowseRecipesState {
	return &BrowseRecipesState{NewStateData(parent), nil}
}

func (s *BrowseRecipesState) GetOptions() Options {
	options := NewOptions(NewLine("Browsing Items"))
	recipes, err := s.Services.Recipes.ReadAll()
	if err != nil {
		panic(err)
	}
	s.recipes = recipes
	for i, recipe := range s.recipes {
		options = options.Add(
			NewOption(fmt.Sprint(i+1), "%s", recipe.RecipeName),
		)
	}
	return options.Add(
		NewOption("A", "Add a recipe"),
		NewOption("B", "Back"),
	)
}
func (s *BrowseRecipesState) Run(option string) State {
	switch option {
	case "A":
		return NewAddRecipeState(s, GetInput("Name"))
	case "B":
		return s.Parent
	default:
		i, _ := strconv.Atoi(option)
		return NewEditRecipeState(s, s.recipes[i-1])
	}
}
