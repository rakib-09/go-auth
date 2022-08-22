package cmd

import (
	"fmt"
	"github.com/rakib-09/go-auth/config"
	"github.com/rakib-09/go-auth/db"
	"github.com/spf13/cobra"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Use: "github.com/rakib-09/go-auth",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

func Execute() {
	config.LoadConfig()
	db.ConnectDB()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
