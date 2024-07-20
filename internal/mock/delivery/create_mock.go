package delivery

import (
	"WbTest/internal/mock/model"
	"encoding/json"
	"net/http"
)

// CreateMock обрабатывает запрос для создания нового мока.
func (d *MockDelivery) CreateMock(w http.ResponseWriter, r *http.Request) {
	var mock model.Mock
	if err := json.NewDecoder(r.Body).Decode(&mock); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := d.service.CreateMock(r.Context(), &mock)
	if err != nil {
		http.Error(w, "Failed to create mock", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
