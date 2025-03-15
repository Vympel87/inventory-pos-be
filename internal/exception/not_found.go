package exception

func NewNotFoundException(message string, errorCode int) *HttpException {
	return NewHttpException(message, errorCode, 404, nil)
}
