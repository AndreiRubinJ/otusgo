package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw13_http/internal/models"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	response := models.GetData{
		Message:     "Hello, this is a GET response!",
		CurrentDate: time.Now().Format(time.RFC3339), // Use RFC3339 for a standard date-time format
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
	fmt.Println("response: ", response)
}
