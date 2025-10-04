package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var req previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validation
	if len(req.UserID) == 0 {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: "created"}

	writeJSON(w, http.StatusCreated, response)
}
