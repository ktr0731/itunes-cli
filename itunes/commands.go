package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/everdev/mack"
	pipeline "github.com/mattn/go-pipeline"
	"github.com/urfave/cli"
)

type SelectType string

var (
	SelectTypeTracks   SelectType = "tracks"
	SelectTypePlayList SelectType = "playlists"
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
		Action:  prev,
	},
	{
		Name:    "back",
		Aliases: []string{"b"},
		Usage:   "Replay current music or play previous music",
		Action:  back,
	},
	{
		Name:      "vol",
		Aliases:   []string{"v"},
		Usage:     "Change volume with an argument (0 - 100)",
		Action:    vol,
		ArgsUsage: "volume",
	},
	{
		Name:    "find",
		Aliases: []string{"f"},
		Usage:   "Find a music (or playlist) by fuzzy search apps",
		Action:  find,
	},
	{
		Name:    "list",
		Aliases: []string{"l", "list", "ls"},
		Usage:   "List all music names of iTunes",
		Action:  list,
	},
}

func play(c *cli.Context) error {
	if c.NArg() > 1 {
		cli.ShowCommandHelp(c, "play")
		return fmt.Errorf("\ninvalid arguments number")
	}

	var err error
	if c.NArg() == 1 {
		_, err = mack.Tell("iTunes", `play track "`+c.Args()[0]+`"`)
	} else {
		_, err = mack.Tell("iTunes", "play")
	}

	if err != nil {
		return fmt.Errorf("cannot play music: %s", err)
	}

	return nil
}

func pause(c *cli.Context) error {
	if _, err := mack.Tell("iTunes", "pause"); err != nil {
		return fmt.Errorf("cannot pause current music: %s", err)
	}

	return nil
}

func next(c *cli.Context) error {
	if _, err := mack.Tell("iTunes", "next track"); err != nil {
		return fmt.Errorf("cannot play next music: %s", err)
	}

	return nil
}

func prev(c *cli.Context) error {
	if _, err := mack.Tell("iTunes", "previous track"); err != nil {
		return fmt.Errorf("cannot play previous music: %s", err)
	}

	return nil
}

func back(c *cli.Context) error {
	if _, err := mack.Tell("iTunes", "back track"); err != nil {
		return fmt.Errorf("cannot back music: %s", err)
	}

	return nil
}

func vol(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "vol")
		return fmt.Errorf("\ninvalid arguments number")
	}

	n, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return fmt.Errorf("cannot convert argument to number: %s", err)
	}

	if n < 0 || n > 100 {
		return fmt.Errorf("invalid range: %d", n)
	}

	if _, err = mack.Tell("iTunes", fmt.Sprintf("set sound volume to %d", n)); err != nil {
		return fmt.Errorf("cannot change volume: %s", err)
	}

	return nil
}

func find(c *cli.Context) error {
	if c.NArg() > 1 {
		cli.ShowCommandHelp(c, "find")
		return fmt.Errorf("\ninvalid arguments number")
	}

	selectType := SelectTypeTracks
	if c.NArg() != 0 {
		selectType = getSelectType(c.Args()[0])
	}

	list, err := listMusics(selectType)
	if err != nil {
		return fmt.Errorf("cannot get music list: %s", err)
	}

	out, err := pipeline.Output(
		[]string{"echo", list},
		[]string{os.Getenv(fuzzy)},
	)

	if err != nil {
		// Ctrl+C
		if strings.Contains("exit status 130", err.Error()) {
			return nil
		}
		return fmt.Errorf("cannot start fuzzy-search: %s", err)
	}

	if strings.TrimSpace(string(out)) == "" {
		return fmt.Errorf("cannot select empty string")
	}

	if _, err = mack.Tell("iTunes", fmt.Sprintf(`play %s "%s"`, string(selectType), strings.Replace(strings.TrimSpace(string(out)), `"`, `\"`, -1))); err != nil {
		return fmt.Errorf("cannot play music: %s", err)
	}

	return nil
}

func list(c *cli.Context) error {
	selectType := SelectTypeTracks
	if c.NArg() != 0 {
		selectType = getSelectType(c.Args()[0])
	}
	fmt.Println(listMusics(selectType))
	return nil
}

func listMusics(selectType SelectType) (string, error) {
	script := `
	set tx to ""
	repeat with t in ` + string(selectType) + `
		set tx to tx & (name of t) & "\\n"
	end
	tx
	`
	mackOut, err := mack.Tell("iTunes", script)
	if err != nil {
		return "", err
	}
	list := strings.TrimSpace(strings.Join(strings.Split(mackOut, "\\n"), "\n"))
	return list, nil
}

func getSelectType(t string) SelectType {
	switch t {
	case "music", "track":
		return SelectTypeTracks
	case "plist":
		return SelectTypePlayList
	default:
		return SelectTypeTracks
	}
}
