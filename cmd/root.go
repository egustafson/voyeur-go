package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "voyeur <sub-command>",
	Short:             "uber executable for voyeur",
	PersistentPreRunE: initAppHook,
}

// flags initialized in flags.go

func Execute() error {
	return rootCmd.Execute()
}

func initAppHook(cmd *cobra.Command, args []string) (err error) {

	//
	// app initialization
	//

	return nil
}
