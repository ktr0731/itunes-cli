package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const fuzzy = "ITUNES_CLI_FUZZY_TOOL"

func init() {
	if os.Getenv(fuzzy) == "" {
		fmt.Fprintln(os.Stderr, "please set environment variable: $"+fuzzy)
		os.Exit(1)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "iTunes CLI"
	app.Usage = "Command line interface for control iTunes"
	app.Version = "0.1.0"
	app.Commands = commands
	app.Run(os.Args)
}
