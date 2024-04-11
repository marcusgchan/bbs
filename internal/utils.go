package internal

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(component templ.Component, c echo.Context) error {
	return component.Render(c.Request().Context(), c.Response())
}

func FromHTMX(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") != ""
}
