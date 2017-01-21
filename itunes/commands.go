package main

import "github.com/urfave/cli"

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
		Action:  nil,
	},
	{
		Name:    "next",
		Aliases: []string{"n", "ne"},
		Usage:   "Play next music",
		Action:  nil,
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
		Action:  nil,
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
