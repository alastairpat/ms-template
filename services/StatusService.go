package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var metadataFileName = "metadata.json"

type StatusResponse struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSHA string `json:"lastcommitsha"`
}

type MetadataFile struct {
	Name string
	Description string
	Version string
}

func BuildStatus() (map[string]interface{}, error) {
	commitSha := getCommitSHA()

	metadata, e := getServiceMetadata()

	if e != nil {
		return nil, e
	}

	status := &StatusResponse{
		Version:       metadata.Version,
		Description:   metadata.Description,
		LastCommitSHA: commitSha,
	}
	
	status.LastCommitSHA = commitSha
	statusWrapper := make(map[string]interface{})
	statusWrapper[metadata.Name] = [1]StatusResponse{*status}
	return statusWrapper, nil
}

func getServiceMetadata() (*MetadataFile, error) {
	fileContents, err := ioutil.ReadFile(metadataFileName)
	var metadata MetadataFile
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata.json: %v", err)
	}
	if len(fileContents) == 0 {
		return nil, fmt.Errorf("failed to read metadata.json: zero bytes read")
	}
	err = json.Unmarshal(fileContents, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata.json: %v", err)
	}

	return &metadata, nil
}

func getCommitSHA() string {
	commitSha := os.Getenv("BUILD_SHA")
	if commitSha == "" {
		commitSha = "HEAD"
	}
	return commitSha
}
