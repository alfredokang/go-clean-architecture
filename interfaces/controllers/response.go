package controllers

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseMessageError struct {
	Error string `json:"error"`
}

func NewResponseMessage(msg string) ResponseMessage {
	return ResponseMessage{
		msg,
	}
}

func NewResponseMessageError(msg string) ResponseMessageError {
	return ResponseMessageError{
		msg,
	}
}
