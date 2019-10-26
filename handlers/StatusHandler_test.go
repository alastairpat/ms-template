package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetStatusIndex(t *testing.T) {
	_ = os.Setenv("BUILD_SHA", "a1b2c3d")
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	err := GetStatusIndex(c)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"example-service\":[{\"version\":\"1.0.1\",\"description\":\"An Example Microservice\",\"lastcommitsha\":\"a1b2c3d\"}]}\n", rec.Body.String())
}
