package command

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	jtime "github.com/jiro4989/go-timer/time"
)

func CmdTimer(c *cli.Context) {
	s := c.Int("seconds") // 秒
	m := c.Int("minutes") // 分
	h := c.Int("hours")   // 時

	sec := s
	sec += m * 60
	sec += h * 60 * 60

	n := time.Now()
	endTime := n.Add(time.Duration(sec) * time.Second)
	fmt.Println("    Start Time:", n.Format("15:04:05"))
	fmt.Println("      End Time:", endTime.Format("15:04:05"))

	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)

	for {
		now := time.Now()
		diff := endTime.Sub(now)
		i := int(diff.Seconds())
		if i < 0 {
			break
		}

		t := jtime.TimeHmsString(i)

		var color *color.Color
		switch {
		case i < 60:
			color = red
		case i < 60*3:
			color = yellow
		default:
			color = green
		}

		color.Print("\rRemaining Time: ", t)
		time.Sleep(1 * time.Second)
	}
}
