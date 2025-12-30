package app

import (
	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
)

type AddItemState struct {
	Parent   State
	Services services.Services

	item  schemas.DBItem
	error string
}

func NewAddItemState(
	parent State,
	s services.Services,
	name, unit string,
) *AddItemState {
	return &AddItemState{
		Parent:   parent,
		Services: s,
		item:     schemas.DBItem{ItemName: name, Unit: unit},
	}
}

func (s *AddItemState) GetError() string {
	return s.error
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
