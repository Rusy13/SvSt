package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// GetMock обрабатывает запрос для получения мока по методу и URL.
func (d *MockDelivery) GetMock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	method := vars["method"]
	url := "/" + vars["url"]

	log.Printf("GetMock called with method: %s, url: %s", method, url)

	mock, err := d.service.GetMock(r.Context(), method, url)
	if err != nil {
		log.Printf("Error fetching mock: %v", err)
		http.Error(w, "Mock not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(mock); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
