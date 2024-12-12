package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw13_http/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/get", nil)
	rr := httptest.NewRecorder()
	GetHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	var response models.GetData
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, this is a GET response!", response.Message)
	_, err = time.Parse(time.RFC3339, response.CurrentDate)
	assert.NoError(t, err)
}

func TestGetHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/get", nil)
	rr := httptest.NewRecorder()
	GetHandler(rr, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	assert.Equal(t, "Invalid request method\n", rr.Body.String())
}
