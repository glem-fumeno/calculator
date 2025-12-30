package app

import (
	"fmt"
	"strconv"

	"github.com/glem-fumeno/calculator/schemas"
)

type BrowseItemsState struct {
	*StateData

	items []schemas.DBItem
}

func NewBrowseItemsState(parent State) *BrowseItemsState {
	return &BrowseItemsState{NewStateData(parent), nil}
}

func (s *BrowseItemsState) GetOptions() Options {
	items, err := s.Services.Items.ReadAll()
	if err != nil {
		panic(err)
	}
	s.items = items
	options := NewOptions(NewLine("Browsing Items"))
	for i, item := range items {
		options = options.Add(
			NewOption(fmt.Sprint(i+1), "%s", item.ItemName),
		)
	}
	return options.Add(
		NewOption("A", "Add an item"),
		NewOption("B", "Back"),
	)
}
func (s *BrowseItemsState) Run(option string) State {
	switch option {
	case "A":
		return NewAddItemState(s, GetInput("Name"), GetInput("Unit"))
	case "B":
		return s.Parent
	default:
		i, _ := strconv.Atoi(option)
		return NewEditItemState(s, s.items[i-1])
	}
}
