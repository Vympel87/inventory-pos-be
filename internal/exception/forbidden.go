package exception

func NewForbiddenException(message string, errorCode int, errors interface{}) *HttpException {
	return NewHttpException(message, errorCode, 403, errors)
}
