package app

import (
	"github.com/glem-fumeno/calculator/schemas"
)

type AddItemState struct {
	*StateData

	item  schemas.DBItem
}

func NewAddItemState( parent State, name, unit string) *AddItemState {
	return &AddItemState{
		NewStateData(parent),
		     schemas.DBItem{ItemName: name, Unit: unit},
	}
}

func (s *AddItemState) GetOptions() Options {
	return NewOptions(
		NewLine("Adding an item"),
		NewOption("N", "Name: %s", s.item.ItemName),
		NewOption("U", "Unit: %s", s.item.Unit),
		NewOption("S", "Save and go back"),
		NewOption("B", "Back"),
	)
}
func (s *AddItemState) Run(option string) State {
	switch option {
	case "N":
		s.item.ItemName = GetInput("Name")
	case "U":
		s.item.Unit = GetInput("Unit")
	case "S":
		err := s.Services.Items.Create(s.item)
		if err == nil {
			return s.Parent
		}
		s.error = err.Error()
	case "B":
		return s.Parent
	}
	return s
}
