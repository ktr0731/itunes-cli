package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "iTunes CLI"
	app.Usage = "Command line interface for control iTunes"
	app.Version = "0.1.0"
	app.Commands = commands
	app.Run(os.Args)
}
