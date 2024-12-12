package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw13_http/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestPostHandler_Success(t *testing.T) {
	requestBody := models.GetData{
		Message: "Test message",
	}
	requestBodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewReader(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	PostHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	var response models.GetData
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Received: Test message", response.Message)

	_, err = time.Parse(time.RFC3339, response.CurrentDate)
	assert.NoError(t, err)
}

func TestPostHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/post", nil)
	rr := httptest.NewRecorder()
	PostHandler(rr, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	assert.Equal(t, "Invalid request method\n", rr.Body.String())
}

func TestPostHandler_InvalidJSON(t *testing.T) {
	invalidJSON := `{"Message":`
	req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewReader([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	PostHandler(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid JSON format\n", rr.Body.String())
}
