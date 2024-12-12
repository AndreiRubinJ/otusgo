package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var originalDir = getWorkingDir()

func TestGetServerUrl_Success(t *testing.T) {
	setWorkingDir()
	os.Setenv("REQUEST_SERVER_URL", "localhost")
	os.Setenv("SERVER_PORT", "8080")
	defer os.Clearenv()
	defer unSetDirectory(originalDir)
	serverURL := getServerURL()
	assert.Equal(t, "localhost:8080", serverURL)
}

func TestStartServer_RouterConfiguration(t *testing.T) {
	router := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "/getUser", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
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
