package recipes

func (q *Queries) Delete(recipeName string) error {
	const query = `
		DELETE FROM recipe_ WHERE recipe_name_ = ?
	`
	_, err := q.connection.Exec(query, recipeName)
	return err
}
