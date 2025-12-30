package app

import (
	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
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
func (s *EditItemState) GetOptions() Options {
	return NewOptions(
		NewLine("Editing %s", s.itemName),
		NewOption("N", "Name: %s", s.item.ItemName),
		NewOption("U", "Unit: %s", s.item.Unit),
		NewOption("D", "Delete"),
		NewOption("S", "Save and go back"),
		NewOption("B", "Back"),
	)
}
func (s *EditItemState) Run(option string) State {
	switch option {
	case "N":
		s.item.ItemName = GetInput("Name")
	case "U":
		s.item.Unit = GetInput("Unit")
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
