package items

func (q *Queries) Delete(itemName string) error {
	const query = `
		DELETE FROM item_ WHERE item_name_ = ?
	`
	_, err := q.connection.Exec(query, itemName)
	return err
}
