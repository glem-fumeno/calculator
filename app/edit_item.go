package app

import (
	"fmt"

	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
	"github.com/glem-fumeno/calculator/tui"
)

type EditItemState struct {
	Parent   State
	Services services.Services

	itemName string
	item     schemas.DBItem
	error    string
}

func NewEditItemState(
	parent State,
	s services.Services,
	item schemas.DBItem,
) *EditItemState {
	return &EditItemState{
		Parent:   parent,
		Services: s,
		itemName: item.ItemName,
		item:     item,
	}
}

func (s *EditItemState) GetError() string {
	return s.error
}
func (s *EditItemState) GetTitle() string {
	return fmt.Sprintf("Editing %s", s.itemName)
}
func (s *EditItemState) GetOptions() []tui.Option {
	return []tui.Option{
		tui.NewOption("N", fmt.Sprintf("Name: %s", s.item.ItemName)),
		tui.NewOption("U", fmt.Sprintf("Unit: %s", s.item.Unit)),
		tui.NewOption("D", "Delete"),
		tui.NewOption("S", "Save and go back"),
		tui.NewOption("B", "Back"),
	}
}
func (s *EditItemState) Run(option string) State {
	switch option {
	case "N":
		s.item.ItemName = tui.GetInput("Name")
	case "U":
		s.item.Unit = tui.GetInput("Unit")
	case "D":
		err := s.Services.Items.Delete(s.itemName)
		if err == nil {
			return s.Parent
		}
		s.error = err.Error()
	case "S":
		err := s.Services.Items.Update(s.itemName, s.item)
		if err == nil {
			return s.Parent
		}
		s.error = err.Error()
	case "B":
		return s.Parent
	}
	return s
}
