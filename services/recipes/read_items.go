package recipes

import (
	"github.com/glem-fumeno/calculator/queries"
	"github.com/glem-fumeno/calculator/schemas"
)

func (s *Services) ReadItems(name string) ([]schemas.DBRecipeItem, []schemas.DBRecipeItem, error) {
	queries, err := queries.NewQueries(s.database)
	if err != nil {
		return nil, nil, err
	}
	defer queries.Rollback()
	items, err := queries.Recipes.SelectItems(name)
	if err != nil {
		return nil, nil, err
	}
	ingredients := make([]schemas.DBRecipeItem, 0)
	products := make([]schemas.DBRecipeItem, 0)
	for _, item := range items {
		switch item.ItemType {
		case schemas.ItemTypeIngredient:
			ingredients = append(ingredients, item)
		case schemas.ItemTypeProduct:
			products = append(products, item)
		}
	}
	return ingredients, products, nil
}
