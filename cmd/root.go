package cmd

import (
	"github.com/spf13/cobra"
	"go-auth/config"
	"go-auth/conn"
)

var (
	RootCmd = &cobra.Command{
		Use: "go-auth",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

func Execute() {
	config.LoadConfig()
	conn.ConnectDB()
}
