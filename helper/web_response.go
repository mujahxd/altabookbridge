package helper

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	jsonResponse := Response{
		Data:    data,
		Message: message,
		Code:    code,
		Status:  status,
	}

	return jsonResponse

}
