package main

import (
	"github.com/spf13/cobra"
	"github.com/werks/voyeur-go/daemon"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start as a daemon",
	RunE:  doDaemon,
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func doDaemon(cmd *cobra.Command, args []string) error {

	err := daemon.Run()
	return err
}
