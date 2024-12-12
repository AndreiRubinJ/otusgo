package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AndreiRubinJ/otusgo/hw13_http/internal/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data models.GetData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	response := models.GetData{
		Message:     "Received: " + data.Message,
		CurrentDate: time.Now().Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
	fmt.Println("response: ", response)
}
