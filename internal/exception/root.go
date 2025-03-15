package exception

type HttpException struct {
	Message    string      `json:"message"`
	ErrorCode  int         `json:"errorCode"`
	StatusCode int         `json:"statusCode"`
	Errors     interface{} `json:"errors,omitempty"`
}

func (e *HttpException) Error() string {
	return e.Message
}

func NewHttpException(message string, errorCode, statusCode int, errors interface{}) *HttpException {
	return &HttpException{
		Message:    message,
		ErrorCode:  errorCode,
		StatusCode: statusCode,
		Errors:     errors,
	}
}
