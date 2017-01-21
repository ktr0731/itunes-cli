package main

import (
	"fmt"

	"github.com/everdev/mack"
	"github.com/urfave/cli"
)

func play(c *cli.Context) error {
	err := mack.Tell("iTunes", "play")
	if err != nil {
		return fmt.Errorf("cannot play music: %s", err)
	}

	return nil
}
