package utils

import (
	"encoding/json"
	"net/http"

	"github.com/stadio-app/stadio-backend/types"
)

func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		json.NewEncoder(w).Encode(
			types.Response{
				Message: "could not parse data",
			},
		)
	}
}

func ErrorResponse(w http.ResponseWriter, status int, errors ...string) {
	JsonResponse(w, status, types.Errors{
		Errors: errors,
	})
}

// types.Data json response with status code 200
func DataResponse(w http.ResponseWriter, data interface{}) {
	JsonResponse(w, http.StatusOK, types.Result{
		Data: data,
	})
}

// types.Data json response with status code 201
func DataResponseCreated(w http.ResponseWriter, data interface{}) {
	JsonResponse(w, http.StatusCreated, types.Result{
		Data: data,
	})
}
