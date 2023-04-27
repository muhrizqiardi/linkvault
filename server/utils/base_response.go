package utils

import (
	"encoding/json"
	"net/http"
)

type BaseResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func CreateBaseResponse[T any](success bool, message string, data T) BaseResponse[T] {
	return BaseResponse[T]{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func BaseResponseWriter[T any](w http.ResponseWriter, statusCode int, success bool, message string, data T) {
	baseResponse := CreateBaseResponse(success, message, data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(baseResponse)
	return
}
