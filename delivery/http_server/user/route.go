package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	e.GET("/export-data", h.GetAll)
}
