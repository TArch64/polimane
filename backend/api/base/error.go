package base

type CustomErrorData map[string]interface{}

type CustomError struct {
	Code    int
	Message string
	Data    CustomErrorData
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(code int, message string, data CustomErrorData) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewReasonedError(code int, reason string) *CustomError {
	return NewCustomError(code, reason, CustomErrorData{"reason": reason})
}
