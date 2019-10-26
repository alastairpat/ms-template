package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type statusResponse struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSHA string `json:"lastcommitsha"`
}

func GetStatusIndex(c echo.Context) error {
	statusWrapper := buildStatus()

	return c.JSON(http.StatusOK, statusWrapper)
}

func buildStatus() map[string]interface{} {
	commitSha := os.Getenv("BUILD_SHA")
	if commitSha == "" {
		commitSha = "HEAD"
	}

	status := &statusResponse{
		Version:       "v",
		Description:   "d",
		LastCommitSHA: commitSha,
	}
	statusWrapper := make(map[string]interface{})
	statusWrapper["alastair"] = [1]statusResponse{*status}
	return statusWrapper
}
