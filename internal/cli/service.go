package cli

import "github.com/urfave/cli"

type App struct {
	*cli.App
}

func NewApp(opts ...func(*App)) *App {
	app_ := cli.NewApp()
	app := &App{
		app_,
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}