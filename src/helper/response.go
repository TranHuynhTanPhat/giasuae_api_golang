package helper

// import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObjec struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildResponseError(message string, err string, data interface{}) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   err,
		Data:    data,
	}
	return res
}
