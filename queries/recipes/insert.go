package recipes

import (
	"github.com/glem-fumeno/calculator/schemas"
)

func (q *Queries) Insert(recipe schemas.DBRecipe) error {
	const query = `
		INSERT INTO recipe_ (recipe_name_)
		VALUES (?)
	`
	_, err := q.connection.Exec(query, recipe.RecipeName)
	return err
}
