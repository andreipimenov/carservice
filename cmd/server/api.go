package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreipimenov/carservice/model"
)

// APIError represents common api error response.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// WriteAPIError writes APIError into http.ResponseWriter as JSON.
func WriteAPIError(w http.ResponseWriter, code string, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	j, _ := json.Marshal(&APIError{
		Code:    code,
		Message: message,
	})
	w.Write(j)
}

// NotFoundHandler for all unsupported API endpoints.
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteAPIError(w, "NotFound", "Unsupported API endpoint", http.StatusNotFound)
	})
}

// Pinghandler for health-checking - /api/ping.
func PingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	})
}

// CarsHandler - GET /api/cars?serialNumber=<uint64>
func CarsHandler(carInteractor model.CarInteractor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			WriteAPIError(w, "BadMethod", "Request method must being GET", http.StatusBadRequest)
			return
		}

		r.ParseForm()
		serialNumber, err := strconv.ParseUint(r.FormValue("serialNumber"), 10, 64)
		if err != nil {
			WriteAPIError(w, "BadRequest", "Parameter serialNumber must being correct uint64", http.StatusBadRequest)
			return
		}

		car, err := carInteractor.FindBySerialNumber(serialNumber)
		if err != nil {
			WriteAPIError(w, "NotFound", fmt.Sprintf("Error: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(car.JSON())
	})
}
