package command

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	jtime "github.com/jiro4989/go-timer/time"
)

func CmdStopwatch(c *cli.Context) {
	fmt.Println("  Start Time:", time.Now().Format("03:04:06"))

	startTime := time.Now()
	for {
		now := time.Now()
		diff := int(now.Sub(startTime).Seconds())
		t := jtime.TimeHmsString(diff)
		fmt.Printf("\rElapsed Time: %s ", t)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
