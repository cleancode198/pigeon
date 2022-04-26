package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/volumefi/conductor/app"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "starts the sparrow server",
		// TODO: add config stuff
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("*chirp chirp*")
			return app.Relayer().Start(cmd.Context())
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
