package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DemoHandler interface {
	GetHello(c echo.Context) error
}

type demoHandler struct {}

func NewDemoHandler() DemoHandler {
	return &demoHandler{}
}

func (dh demoHandler) GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Hoge Hoge")
}
