package utils

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Error error
	Type  int
}

type Error struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, apiError *ApiError) {
	w.WriteHeader(apiError.Type)
	json.NewEncoder(w).Encode(Error{
		Error: apiError.Error.Error(),
	})
}
