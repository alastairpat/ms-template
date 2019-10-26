package services

import (
	"os"
	"reflect"
	"testing"
)

func TestBuildStatus(t *testing.T) {
	validMetadataNoSHA := make(map[string]interface{})

	validMetadataNoSHA["example-service"] = [1]StatusResponse{{
		Version:       "1.0.1",
		Description:   "An Example Microservice",
		LastCommitSHA: "HEAD",
	}}

	validMetadataWithSHA := make(map[string]interface{})

	validMetadataWithSHA["example-service"] = [1]StatusResponse{{
		Version:       "1.0.1",
		Description:   "An Example Microservice",
		LastCommitSHA: "a1b2c3d",
	}}

	tests := []struct {
		name    string
		metadataFileName string
		environValue string
		want    map[string]interface{}
		wantErr bool
	}{
		{"metadata doesn't exist", "doesnt_exist.json", "", nil, true},
		{"metadata does exist, is invalid", "invalid.json", "", nil, true},
		{"metadata does exist, is empty", "empty.json", "", nil, true},
		{"metadata does exist, is valid, has BUILD_SHA set", "valid.json", "a1b2c3d", validMetadataWithSHA, false},
		{"metadata does exist, is valid, no BUILD_SHA set", "valid.json", "HEAD", validMetadataNoSHA, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metadataFileName = "test_data/" + tt.metadataFileName
			_ = os.Setenv("BUILD_SHA", tt.environValue)
			got, err := BuildStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
