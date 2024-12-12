package client

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var originalDir = getWorkingDir()

func setupEnv(envVars map[string]string) func() {
	originalEnv := map[string]string{}
	for key := range envVars {
		originalEnv[key] = os.Getenv(key)
	}
	for key, value := range envVars {
		os.Setenv(key, value)
	}

	return func() {
		for key, value := range originalEnv {
			os.Setenv(key, value)
		}
	}
}

func TestParseConfig_SuccessFromEnvironment(t *testing.T) {
	setWorkingDir()
	envVars := map[string]string{
		"REQUEST_METHOD": "GET",
		"REQUEST_PATH":   "/getUser",
		"REQUEST_DATA":   "",
	}
	cleanup := setupEnv(envVars)
	defer cleanup()
	defer unSetDirectory(originalDir)
	defer resetFlags()

	cfg, err := ParseConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "GET", cfg.Method)
	assert.Equal(t, "http://localhost:8080", cfg.URL)
	assert.Equal(t, "/getUser", cfg.Path)
	assert.Equal(t, "", cfg.Data)
}

func TestConfig_ValidateMissingMethod(t *testing.T) {
	setWorkingDir()
	cfg := &Config{
		Method: "",
		URL:    "http://localhost:8080",
		Path:   "/getUser",
	}

	defer unSetDirectory(originalDir)
	defer resetFlags()
	err := cfg.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported HTTP method")
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

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
