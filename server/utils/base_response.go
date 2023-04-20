package utils

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
