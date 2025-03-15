package exception

func NewUnauthorizedException(message string, errorCode int, errors interface{}) *HttpException {
	return NewHttpException(message, errorCode, 401, errors)
}
