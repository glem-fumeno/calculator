package recipes

import "github.com/glem-fumeno/calculator/schemas"

func (q *Queries) Update(recipeName string, recipe schemas.DBRecipe) error {
	const query = `
		UPDATE recipe_ SET
			recipe_name_ = ?
		WHERE recipe_name_ = ?
	`
	_, err := q.connection.Exec(query, recipe.RecipeName, recipeName)
	return err
}
