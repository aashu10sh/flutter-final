package core

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}
