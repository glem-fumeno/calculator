package app

import (
	"fmt"

	"github.com/glem-fumeno/calculator/schemas"
	"github.com/glem-fumeno/calculator/services"
	"github.com/glem-fumeno/calculator/tui"
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
func (s *AddItemState) GetTitle() string {
	return "Adding an item"
}
func (s *AddItemState) GetOptions() []tui.Option {
	return []tui.Option{
		tui.NewOption("N", fmt.Sprintf("Name: %s", s.item.ItemName)),
		tui.NewOption("U", fmt.Sprintf("Unit: %s", s.item.Unit)),
		tui.NewOption("S", "Save and go back"),
		tui.NewOption("B", "Back"),
	}
}
func (s *AddItemState) Run(option string) State {
	switch option {
	case "N":
		s.item.ItemName = tui.GetInput("Name")
	case "U":
		s.item.Unit = tui.GetInput("Unit")
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
