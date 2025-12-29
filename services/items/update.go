package items

import (
	"errors"

	"github.com/glem-fumeno/calculator/queries"
	"github.com/glem-fumeno/calculator/schemas"
)

func (s *Services) Update(name string, item schemas.DBItem) error {
	if len(item.ItemName) < 1 {
		return errors.New("Item name must not be empty")
	}
	if len(item.Unit) < 1 {
		return errors.New("Unit must not be empty")
	}
	queries, err := queries.NewQueries(s.database)
	if err != nil {
		return err
	}
	defer queries.Rollback()
	err = queries.Items.Update(name, item)
	if err != nil {
		return err
	}
	return queries.Commit()
}
