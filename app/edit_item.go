package app

import (
	"github.com/glem-fumeno/calculator/schemas"
)

type EditItemState struct {
	*StateData

	itemName string
	item     schemas.DBItem
}

func NewEditItemState(parent State, item schemas.DBItem) *EditItemState {
	return &EditItemState{NewStateData(parent), item.ItemName, item}
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
