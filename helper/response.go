package helper

//Response is used for static shape json return
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Data:    data,
	}
	return res
}
