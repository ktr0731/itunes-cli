package main

import (
	"fmt"

	"github.com/everdev/mack"
	"github.com/urfave/cli"
)

func pause(c *cli.Context) error {
	err := mack.Tell("iTunes", "pause")
	if err != nil {
		return fmt.Errorf("cannot pause current music: %s", err)
	}

	return nil
}
