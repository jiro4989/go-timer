package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "go-timer"
	app.Version = "1.0.0"
	app.Author = "jiro4989"
	app.Email = ""
	app.Usage = "Timer with Golang on CLI"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
