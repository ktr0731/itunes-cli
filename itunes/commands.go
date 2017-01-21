package main

import (
	"fmt"

	"github.com/everdev/mack"
	"github.com/urfave/cli"
)

var commands = []cli.Command{
	{
		Name:    "play",
		Aliases: []string{"pl", "start"},
		Usage:   "Play current selected music",
		Action:  play,
	},
	{
		Name:    "pause",
		Aliases: []string{"pa", "stop"},
		Usage:   "Stop current playing music",
		Action:  pause,
	},
	{
		Name:    "next",
		Aliases: []string{"n", "ne"},
		Usage:   "Play next music",
		Action:  next,
	},
	{
		Name:    "prev",
		Aliases: []string{"pr"},
		Usage:   "Play previous music",
		Action:  nil,
	},
	{
		Name:    "back",
		Aliases: []string{"b"},
		Usage:   "Replay current music or play previous music",
		Action:  back,
	},
	{
		Name:    "vol",
		Aliases: []string{"v"},
		Usage:   "Change volume with an argument",
		Action:  nil,
	},
	{
		Name:    "find",
		Aliases: []string{"v"},
		Usage:   "Find a music (or playlist, artist, album) by fuzzy search apps",
		Action:  nil,
	},
}

func play(c *cli.Context) error {
	err := mack.Tell("iTunes", "play")
	if err != nil {
		return fmt.Errorf("cannot play music: %s", err)
	}

	return nil
}

func pause(c *cli.Context) error {
	err := mack.Tell("iTunes", "pause")
	if err != nil {
		return fmt.Errorf("cannot pause current music: %s", err)
	}

	return nil
}

func next(c *cli.Context) error {
	err := mack.Tell("iTunes", "next track")
	if err != nil {
		return fmt.Errorf("cannot play next music: %s", err)
	}

	return nil
}

func back(c *cli.Context) error {
	err := mack.Tell("iTunes", "back track")
	if err != nil {
		return fmt.Errorf("cannot back music: %s", err)
	}

	return nil
}
