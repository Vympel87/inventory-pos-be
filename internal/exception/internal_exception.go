package exception

func NewInternalException(message string, errors interface{}, errorCode int) *HttpException {
	return NewHttpException(message, errorCode, 500, errors)
}
