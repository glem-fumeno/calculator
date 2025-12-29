package items

import (
	"github.com/glem-fumeno/calculator/queries"
)

func (s *Services) Delete(name string) error {
	queries, err := queries.NewQueries(s.database)
	if err != nil {
		return err
	}
	defer queries.Rollback()
	err = queries.Items.Delete(name)
	if err != nil {
		return err
	}
	return queries.Commit()
}
