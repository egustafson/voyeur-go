package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "voyeur <sub-command>",
	Short:             "uber executable for voyeur",
	PersistentPreRunE: initAppHook,
}

// flags initialized in flags.go

func main() {
	// this program uses Cobra (https://github.com/spf13/cobra)
	err := Execute()
	if err != nil {
		// cobra will print an error to stdout/(?)err
		os.Exit(1)
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func initAppHook(cmd *cobra.Command, args []string) (err error) {

	//
	// app initialization
	//

	return nil
}
