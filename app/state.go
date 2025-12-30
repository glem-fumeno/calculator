package app

import "github.com/glem-fumeno/calculator/services"

type State interface {
	GetError() string
	GetServices() services.Services
	GetOptions() Options
	Run(option string) State
}

type StateData struct {
	Parent   State
	Services services.Services
	error    string
}

func NewStateData(parent State) *StateData {
	return &StateData{parent, parent.GetServices(), ""}
}

func (s *StateData) GetError() string {
	return s.error
}
func (s *StateData) GetServices() services.Services {
	return s.Services
}

type App struct {
	State State
}

func (a *App) Run() {
	for {
		option := GetOption(
			a.State.GetError(),
			a.State.GetOptions(),
		)
		a.State = a.State.Run(option)
		if a.State == nil {
			return
		}
	}
}
