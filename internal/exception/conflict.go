package exception

func NewConflictException(message string, errorCode int, errors interface{}) *HttpException {
	return NewHttpException(message, errorCode, 409, errors)
}
