package helpers

import (
	"encoding/json"
	"net/http"
	"othello-backend/models"
	"strconv"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func FetchGameCode(r *http.Request) int {
	otpStr := r.PathValue("id")
	otp, _ := strconv.Atoi(otpStr)
	return otp
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error: msg,
	})
}
