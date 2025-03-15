package exception

func NewBadRequestException(message string, errors interface{}, errorCode int) *HttpException {
	return NewHttpException(message, errorCode, 400, errors)
}
