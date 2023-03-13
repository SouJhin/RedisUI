package main

import (
	"context"
	"fmt"

	"changeme/internal/define"
	"changeme/internal/service"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ConnectionList() H {
	conn, err := service.ConnectionList()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR:" + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": conn,
	}
}

func (a *App) ConnectionCreate(connection *define.Connection) H {
	return M{
		"code": 200,
		"msg":  "connection success",
	}
}
