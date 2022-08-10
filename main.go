package main

import (
	"github.com/spf13/viper"
	"go-auth/cmd"
)

func main() {
	viper.SetConfigFile(".env")
	cmd.Execute()
}
