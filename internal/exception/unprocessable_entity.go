package exception

func NewUnprocessableEntityException(message string, errorCode int, errors interface{}) *HttpException {
	return NewHttpException(message, errorCode, 422, errors)
}
