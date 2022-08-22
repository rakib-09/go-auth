package db

import (
	"github.com/rakib-09/go-auth/domains"
)

var client AuthDatabase

func NewDbClient() domains.Database {
	ConnectDB()
	return &AuthDatabase{}
}

func Client() *AuthDatabase {
	return &client
}
