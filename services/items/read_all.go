package items

import (
	"github.com/glem-fumeno/calculator/queries"
	"github.com/glem-fumeno/calculator/schemas"
)

func (s *Services) ReadAll() ([]schemas.DBItem, error) {
	queries, err := queries.NewQueries(s.database)
	if err != nil {
		return nil, err
	}
	defer queries.Rollback()
	items, err := queries.Items.SelectMany()
	if err != nil {
		return nil, err
	}
	return items, nil
}
