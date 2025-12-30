package recipes

import "github.com/glem-fumeno/calculator/schemas"

func (q *Queries) SelectMany() ([]schemas.DBRecipe, error) {
	const query = `
		SELECT recipe_name_
		FROM recipe_
	`
	rows, err := q.connection.Query(query)
	var result []schemas.DBRecipe = make([]schemas.DBRecipe, 0)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		item := schemas.DBRecipe{}
		rows.Scan(&item.RecipeName)
		result = append(result, item)
	}
	return result, nil
}
