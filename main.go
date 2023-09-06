package main

import (
	"log"

	"github.com/koor-tech/data-control-center/pkg/server"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "server-app",
		Short: "A demo application for starting different types of servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			server.StartHTTPServer()
			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
