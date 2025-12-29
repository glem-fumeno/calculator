package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/glem-fumeno/calculator/queries"
	"github.com/glem-fumeno/calculator/schemas"
)

func browseItems() {
	fmt.Println("Browsing items")
	for {
		options := []Option{}
		items := fetchItems()
		for i, item := range items {
			options = append(
				options,
				Option{
					fmt.Sprint(i + 1),
					fmt.Sprintf("%s (%s)", item.ItemName, item.Unit),
				},
			)
		}
		options = append(options, Option{"A", "Add an item"})
		options = append(options, Option{"B", "Back"})
		option, err := getOption(options)
		if err != nil {
			log.Fatal(err)
		}
		switch option {
		case "A":
			addItem()
		case "B":
			return
		default:
			i, _ := strconv.Atoi(option)
			item := items[i-1]
			editItem(item)
		}
	}
}
func addItem() {
	fmt.Println("Adding an item")
	item := schemas.DBItem{}
	for {
		option, err := getOption([]Option{
			{"N", fmt.Sprintf("Name: %s", item.ItemName)},
			{"U", fmt.Sprintf("Unit: %s", item.Unit)},
			{"S", "Add and go back"},
			{"B", "Back"},
		})
		if err != nil {
			log.Fatal(err)
		}
		switch option {
		case "N":
			name, err := getUserInput("Name")
			if err != nil {
				fmt.Println("Could not get user input")
				continue
			}
			item.ItemName = name
		case "U":
			unit, err := getUserInput("Unit")
			if err != nil {
				fmt.Println("Could not get user input")
				continue
			}
			item.Unit = unit
		case "S":
			fmt.Println(item.ItemName)
			if len(item.ItemName) < 1 {
				fmt.Println("Item name must not be empty")
				continue
			}
			if len(item.Unit) < 1 {
				fmt.Println("Unit must not be empty")
				continue
			}
			createItem(item)
			return
		case "B":
			return
		}
	}
}
func editItem(item schemas.DBItem) {
	itemName := item.ItemName
	for {
		fmt.Printf("Editing %s\n", itemName)
		option, err := getOption([]Option{
			{"N", fmt.Sprintf("Name: %s", item.ItemName)},
			{"U", fmt.Sprintf("Unit: %s", item.Unit)},
			{"D", "Delete"},
			{"S", "Save and go back"},
			{"B", "Back"},
		})
		if err != nil {
			log.Fatal(err)
		}
		switch option {
		case "N":
			name, err := getUserInput("Name")
			if err != nil {
				fmt.Println("Could not get user input")
				continue
			}
			item.ItemName = name
		case "D":
			deleteItem(itemName)
			return
		case "U":
			unit, err := getUserInput("Unit")
			if err != nil {
				fmt.Println("Could not get user input")
				continue
			}
			item.Unit = unit
		case "S":
			if len(item.ItemName) < 1 {
				fmt.Println("Item name must not be empty")
				continue
			}
			if len(item.Unit) < 1 {
				fmt.Println("Unit must not be empty")
				continue
			}
			saveItem(itemName, item)
			return
		case "B":
			return
		}
	}
}

func fetchItems() []schemas.DBItem {
	queries, err := queries.NewQueries("app.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer queries.Close(&err)
	items, err := queries.Items.SelectMany()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return items
}
func saveItem(name string, item schemas.DBItem) {
	queries, err := queries.NewQueries("app.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer queries.Close(&err)
	err = queries.Items.Update(name, item)
	if err != nil {
		fmt.Println(err)
	}
}
func createItem(item schemas.DBItem) {
	queries, err := queries.NewQueries("app.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer queries.Close(&err)
	err = queries.Items.Insert(item)
	if err != nil {
		fmt.Println(err)
	}
}
func deleteItem(name string) {
	queries, err := queries.NewQueries("app.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer queries.Close(&err)
	err = queries.Items.Delete(name)
	if err != nil {
		fmt.Println(err)
	}
}
