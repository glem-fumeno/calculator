package items

import "github.com/glem-fumeno/calculator/schemas"

func (q *Queries) SelectMany() ([]schemas.DBItem, error) {
	const query = `
		SELECT item_name_, unit_
		FROM item_
	`
	rows, err := q.connection.Query(query)
	var result []schemas.DBItem = make([]schemas.DBItem, 0)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		item := schemas.DBItem{
			ItemName: "",
			Unit:     "",
		}
		rows.Scan(&item.ItemName, &item.Unit)
		result = append(result, item)
	}
	return result, nil
}
