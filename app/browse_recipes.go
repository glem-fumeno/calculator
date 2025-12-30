package app

import (
	"fmt"
	"strconv"

	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
)

type BrowseRecipesState struct {
	Parent   State
	Services services.Services

	error   string
	recipes []schemas.DBRecipe
}

func NewBrowseRecipesState(
	parent State,
	s services.Services,
) *BrowseRecipesState {
	return &BrowseRecipesState{Parent: parent, Services: s}
}

func (s *BrowseRecipesState) GetError() string {
	return s.error
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
		return NewAddRecipeState(s, s.Services, GetInput("Name"))
	case "B":
		return s.Parent
	default:
		i, _ := strconv.Atoi(option)
		return NewEditRecipeState(s, s.Services, s.recipes[i-1])
	}
}
