package service_response

type CustomResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

type CustomResponseWithData struct {
	CustomResponse
	Data interface{} `json:"data"`
}

func Error(code int, message string) *CustomResponse {
	return &CustomResponse{false, code, message}
}

func Ok(message string) *CustomResponse {
	return &CustomResponse{true, 0, message}
}

func OkWithData(message string, data interface{}) *CustomResponseWithData {
	return &CustomResponseWithData{CustomResponse{true, 0, message}, data}
}
