package types

type ValidationError struct {
	Error error `json:"validation_error"`
}

type RestErr struct {
	Message string `json:"message"`
}
