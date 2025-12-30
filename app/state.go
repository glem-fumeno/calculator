package app

type State interface {
	GetError() string
	GetOptions() Options
	Run(option string) State
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
