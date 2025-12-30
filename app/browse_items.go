package app

import (
	"fmt"
	"strconv"

	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
)

type BrowseItemsState struct {
	Parent   State
	Services services.Services

	error string
	items []schemas.DBItem
}

func NewBrowseItemsState(
	parent State,
	s services.Services,
) *BrowseItemsState {
	return &BrowseItemsState{Parent: parent, Services: s}
}

func (s *BrowseItemsState) GetError() string {
	return s.error
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
		return NewAddItemState(
			s, s.Services,
			GetInput("Name"), GetInput("Unit"),
		)
	case "B":
		return s.Parent
	default:
		i, _ := strconv.Atoi(option)
		return NewEditItemState(s, s.Services, s.items[i-1])
	}
}
