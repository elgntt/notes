package errors

type BusinessError struct {
	code    string
	message string
}

func (b BusinessError) Error() string {
	return b.message
}

func (b BusinessError) Code() string {
	return b.code
}

func NewBusinessError(message string) BusinessError {
	return BusinessError{
		code:    "BusinessError",
		message: message,
	}
}
