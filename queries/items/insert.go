package items

import "github.com/glem-fumeno/calculator/schemas"

func (q *Queries) Insert(item schemas.DBItem) error {
	const query = `
		INSERT INTO item_ (item_name_, unit_)
		VALUES (?, ?)
	`
	_, err := q.connection.Exec(query, item.ItemName, item.Unit)
	return err
}
