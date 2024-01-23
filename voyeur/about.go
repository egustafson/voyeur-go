package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "about voyeur executable",
	RunE:  doAbout,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func doAbout(cmd *cobra.Command, args []string) error {

	// TODO:  investigate the CLI "output" method
	//

	// output is valid YAML
	fmt.Println("---")
	fmt.Println("voyeur-version: 0.0.0")
	fmt.Println("...")

	return nil
}
