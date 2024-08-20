package userhandler

import (
	userparam "github.com/amiranbari/challenge/param/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) GetAll(c echo.Context) error {
	var req userparam.GetAllRequest

	if bErr := c.Bind(&req); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.userSvc.GetAll(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}
