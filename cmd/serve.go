package cmd

import (
	"github.com/spf13/cobra"
	"go-auth/conn"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	dbClient := conn.DB()
}
