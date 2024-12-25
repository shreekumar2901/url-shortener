package dto

type ErrorResponse struct {
	ErrorMsgs  []string
	StatusCode uint
}

type SuccessResponse struct {
	Message    string
	StatusCode uint
}
