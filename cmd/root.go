package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-auth/config"
	"go-auth/conn"
	"os"
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

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
