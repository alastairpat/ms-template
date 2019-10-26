package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
)

type statusResponse struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSHA string `json:"lastcommitsha"`
}

func GetStatusIndex(c echo.Context) error {
	statusWrapper := buildStatus(c.Logger())

	return c.JSON(http.StatusOK, statusWrapper)
}

func buildStatus(logger echo.Logger) map[string]interface{} {
	commitSha := os.Getenv("BUILD_SHA")
	if commitSha == "" {
		commitSha = "HEAD"
	}

	fileContents, err := ioutil.ReadFile("metadata.json")
	var m map[string]string

	if err != nil || len(fileContents) == 0 {
		logger.Warn("Failed to read metadata.json: ")
		return nil
	}

	err = json.Unmarshal(fileContents, &m)

	if err != nil {
		logger.Warn("Failed to process metadata.json: ")
		return nil
	}

	status := &statusResponse{
		Version:       m["version"],
		Description:   m["description"],
		LastCommitSHA: commitSha,
	}
	statusWrapper := make(map[string]interface{})
	statusWrapper[m["name"]] = [1]statusResponse{*status}
	return statusWrapper
}
