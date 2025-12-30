package recipes

import "github.com/glem-fumeno/calculator/schemas"

func (q *Queries) SelectItems(recipeName string) ([]schemas.DBRecipeItem, error) {
	const query = `
		SELECT recipe_name_, item_name_, item_type_, quantity_
		FROM recipe_item_
		WHERE recipe_name_ = ?
	`
	rows, err := q.connection.Query(query, recipeName)
	var result []schemas.DBRecipeItem = make([]schemas.DBRecipeItem, 0)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		item := schemas.DBRecipeItem{}
		rows.Scan(&item.RecipeName, &item.ItemName, &item.ItemType, &item.Quantity)
		result = append(result, item)
	}
	return result, nil
}
