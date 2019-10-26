package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handles the route GET '/'
func GetRootIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
