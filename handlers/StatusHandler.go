package handlers

import (
	"github.com/alastairpat/ms-template/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetStatusIndex(c echo.Context) error {
	statusWrapper, err := services.BuildStatus()

	if err != nil {
		c.Logger().Errorf("failed to determine service metadata: %v", err)
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, statusWrapper)
}
