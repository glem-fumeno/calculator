package app

import (
	"fmt"
	"strconv"

	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
	"github.com/glem-fumeno/calculator/tui"
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
func (s *BrowseItemsState) GetTitle() string {
	return "Browsing Items"
}
func (s *BrowseItemsState) GetOptions() []tui.Option {
	options := []tui.Option{}
	items, err := s.Services.Items.ReadAll()
	if err != nil {
		panic(err)
	}
	s.items = items
	for i, item := range items {
		options = append(
			options,
			tui.NewOption(
				fmt.Sprint(i+1),
				fmt.Sprintf("%s", item.ItemName),
			),
		)
	}
	return append(
		options,
		tui.NewOption("A", "Add an item"),
		tui.NewOption("B", "Back"),
	)
}
func (s *BrowseItemsState) Run(option string) State {
	switch option {
	case "A":
		return NewAddItemState(
			s, s.Services,
			tui.GetInput("Name"), tui.GetInput("Unit"),
		)
	case "B":
		return s.Parent
	default:
		i, _ := strconv.Atoi(option)
		return NewEditItemState(s, s.Services, s.items[i-1])
	}
}
