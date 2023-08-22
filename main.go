package main

import (
	"fmt"
	"github.com/koor-tech/data-control-center/pkg/server"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var serverType string

	var rootCmd = &cobra.Command{
		Use:   "server-app",
		Short: "A demo application for starting different types of servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			if serverType == "http" {
				fmt.Println("Starting HTTP server...")
				server.StartHTTPServer()
			} else if serverType == "grpc" {
				fmt.Println("Starting gRPC server...")
				server.StartGRPCServer()
			} else {
				return fmt.Errorf("invalid server type: %s", serverType)
			}
			return nil
		},
	}

	rootCmd.Flags().StringVar(&serverType, "server", "", "Choose server type: http or grpc (required)")
	err := rootCmd.MarkFlagRequired("server")
	if err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
