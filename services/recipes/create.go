package recipes

import (
	"errors"

	"github.com/glem-fumeno/calculator/queries"
	"github.com/glem-fumeno/calculator/schemas"
)

func (s *Services) Create(recipe schemas.DBRecipe) error {
	if len(recipe.RecipeName) < 1 {
		return errors.New("Recipe name must not be empty")
	}
	queries, err := queries.NewQueries(s.database)
	if err != nil {
		return err
	}
	defer queries.Rollback()
	err = queries.Recipes.Insert(recipe)
	if err != nil {
		return err
	}
	return queries.Commit()
}
