package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jiro4989/go-timer/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "timer",
		Usage:  "Timer",
		Action: command.CmdTimer,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "seconds, s",
				Value: 0,
				Usage: "",
			},
			cli.IntFlag{
				Name:  "minutes, m",
				Value: 0,
				Usage: "",
			},
			cli.IntFlag{
				Name:  "hours",
				Value: 0,
				Usage: "",
			},
		},
	},
	{
		Name:   "stopwatch",
		Usage:  "Stopwatch",
		Action: command.CmdStopwatch,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "alerm",
		Usage:  "Alertm",
		Action: command.CmdAlerm,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "time,t",
				Value: "",
				Usage: "",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
