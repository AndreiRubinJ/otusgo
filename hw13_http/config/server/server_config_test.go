package server

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var originalDir = getWorkingDir()

func TestParseConfig_WithFlags(t *testing.T) {
	setWorkingDir()
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	defer unSetDirectory(originalDir)
	os.Args = []string{"cmd", "-url=localhost", "-port=8081"}

	config, err := ParseConfig()
	assert.NoError(t, err, "ParseConfig should not return an error")

	// Assertions
	assert.Equal(t, "localhost", config.URL, "URL should be 'localhost'")
	assert.Equal(t, "8081", config.Port, "Port should be '8081'")
}

func TestValidate_ValidConfig(t *testing.T) {
	config := &Config{
		URL:  "http://example.com",
		Port: "8080",
	}

	err := config.Validate()
	assert.NoError(t, err, "Validate should not return an error for a valid config")
}

func TestValidate_InvalidConfig(t *testing.T) {
	config := &Config{
		URL:  "",
		Port: "",
	}

	err := config.Validate()
	assert.Error(t, err, "Validate should return an error for an invalid config")
	assert.Equal(t, "unsupported empty port  or url: ", err.Error(), "Error message should match")
}

func getWorkingDir() string {
	originalDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}
	return originalDir
}
func unSetDirectory(directory string) {
	err := os.Chdir(directory)
	if err != nil {
		fmt.Println("Error restoring the original working directory:", err)
		return
	}
}
func setWorkingDir() {
	err := os.Chdir("../../")
	if err != nil {
		fmt.Println("Error setting new working directory:", err)
		return
	}
}
