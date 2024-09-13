package main

import (
	"fmt"
	"os"
	"timestream-simple-cli/cmd"
)

func main() {
	cmd := cmd.NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
