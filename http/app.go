package http

import (
	"github.com/labstack/echo/v4"
)

type App struct {
	Storage    Storage
	HttpServer *echo.Echo
}

func NewApp(storage Storage) *App {
	return &App{Storage: storage}
}

func (a *App) RunHttp(port string) {
	a.HttpServer = echo.New()

	handler := NewHandler(a.Storage)
	a.HttpServer.GET("/:key", handler.GetValue)
	a.HttpServer.POST("/:key", handler.SetValue)
	a.HttpServer.DELETE("/:key", handler.DeleteValue)

	a.HttpServer.Start(port)
}
