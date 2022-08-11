package controllers

import "go-auth/domains"

type UserController struct {
	authService domains.AuthUseCase
}
