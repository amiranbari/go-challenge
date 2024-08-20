package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	e.Static("/exports", "exports")
	e.GET("/export-data", h.GetAll)
}
