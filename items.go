package main

import (
	"fmt"
	"strconv"

	"github.com/glem-fumeno/calculator/schemas"
)

func browseItems() {
	error := ""
	for {
		options := []Option{}
		items, err := service.Items.ReadAll()
		if err != nil {
			panic(err)
		}
		for i, item := range items {
			options = append(
				options,
				Option{
					fmt.Sprint(i + 1),
					fmt.Sprintf("%s", item.ItemName),
				},
			)
		}
		option := getOption(
			"Browsing items",
			error,
			append(
				options,
				Option{"A", "Add an item"},
				Option{"B", "Back"},
			),
		)
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
	item := schemas.DBItem{}
	error := ""
	item.ItemName = getUserInput("Name")
	item.Unit = getUserInput("Unit")
	for {
		option := getOption(
			"Adding an item",
			error,
			[]Option{
				{"N", fmt.Sprintf("Name: %s", item.ItemName)},
				{"U", fmt.Sprintf("Unit: %s", item.Unit)},
				{"S", "Save and go back"},
				{"B", "Back"},
			})
		switch option {
		case "N":
			item.ItemName = getUserInput("Name")
		case "U":
			item.Unit = getUserInput("Unit")
		case "S":
			err := service.Items.Create(item)
			if err != nil {
				error = err.Error()
				continue
			}
			return
		case "B":
			return
		}
	}
}
func editItem(item schemas.DBItem) {
	itemName := item.ItemName
	error := ""
	for {
		option := getOption(
			fmt.Sprintf("Editing %s", itemName),
			error,
			[]Option{
				{"N", fmt.Sprintf("Name: %s", item.ItemName)},
				{"U", fmt.Sprintf("Unit: %s", item.Unit)},
				{"D", "Delete"},
				{"S", "Save and go back"},
				{"B", "Back"},
			})
		switch option {
		case "N":
			item.ItemName = getUserInput("Name")
		case "U":
			item.Unit = getUserInput("Unit")
		case "D":
			err := service.Items.Delete(itemName)
			if err != nil {
				error = err.Error()
				continue
			}
			return
		case "S":
			err := service.Items.Update(itemName, item)
			if err != nil {
				error = err.Error()
				continue
			}
			return
		case "B":
			return
		}
	}
}
