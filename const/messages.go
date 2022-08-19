package _const

import "errors"

var (
	SomethingWentWrong  = errors.New("something went wrong")
	InvalidCreds        = errors.New("invalid credentials")
	InvalidAccessToken  = errors.New("couldn't generate access token")
	SuccessfullyCreated = "successfully created!"
	SuccessfullyUpdated = "successfully updated!"
)
