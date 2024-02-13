package main

import (
	"os"

	"github.com/werks/voyeur-go/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		// cobra will print an error to stdout/(?)err
		os.Exit(1)
	}
}
