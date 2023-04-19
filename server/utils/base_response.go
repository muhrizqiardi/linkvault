package utils

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func CreateBaseResponse(success bool, message string, data any) BaseResponse {
	return BaseResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}
