package delivery

import (
	"encoding/json"
	"net/http"
)

// GetAllMocks обрабатывает запрос для получения всех моков.
func (d *MockDelivery) GetAllMocks(w http.ResponseWriter, r *http.Request) {
	mocks, err := d.service.GetAllMocks(r.Context())
	if err != nil {
		http.Error(w, "Failed to get mocks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(mocks); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
