package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterHandler(e *echo.Echo) {
	demoHandler := NewDemoHandler()
	e.GET("/hello", demoHandler.GetHello)
}
