package main

import (
	"log"

	"github.com/koor-tech/data-control-center/pkg/server"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "data-control-center",
		Short: "Nice dashboard and controls for managing Rook Ceph data storage.",
		RunE: func(cmd *cobra.Command, args []string) error {
			server.StartHTTPServer()
			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
