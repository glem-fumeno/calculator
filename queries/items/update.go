package items

import "github.com/glem-fumeno/calculator/schemas"

func (q *Queries) Update(itemName string, item schemas.DBItem) error {
	const query = `
		UPDATE item_ SET
			item_name_ = ?,
			unit_ = ?
		WHERE item_name_ = ?
	`
	_, err := q.connection.Exec(query, item.ItemName, item.Unit, itemName)
	return err
}
