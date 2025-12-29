package app

import "github.com/glem-fumeno/calculator/tui"

type State interface {
	GetError() string
	GetTitle() string
	GetOptions() []tui.Option
	Run(option string) State
}

type App struct {
	State State
}

func (a *App) Run() {
	for {
		option := tui.GetOption(
			a.State.GetTitle(),
			a.State.GetError(),
			a.State.GetOptions(),
		)
		a.State = a.State.Run(option)
		if a.State == nil {
			return
		}
	}
}
