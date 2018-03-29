package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	jtime "github.com/jiro4989/go-timer/time"
)

func CmdAlerm(c *cli.Context) (err error) {
	t := c.String("t")
	tl := len(strings.Split(t, ":"))

	switch {
	case tl < 2 || 4 < tl:
		return
	case tl == 2:
		// HH:MM:00
		t += ":00"
	case tl == 3:
		// 何もしない
	default:
		return
	}

	n := time.Now()
	sec, err := jtime.TimeSec(t)
	y, m, d := n.Date()
	h, mm, s := jtime.TimeHms(sec)
	endTime := time.Date(y, m, d, h, mm, s, 0, time.Local)

	fmt.Println("    Start Time:", n.Format("15:04:05"))
	fmt.Println("      End Time:", endTime.Format("15:04:05"))

	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)

	for {
		now := time.Now()
		diff := int(endTime.Sub(now).Seconds())
		if diff < 0 {
			break
		}

		var color *color.Color
		switch {
		case diff < 60:
			color = red
		case diff < 60*3:
			color = yellow
		default:
			color = green
		}

		t := jtime.TimeHmsString(diff)
		color.Print("\rRemaining Time: ", t, " ")
		time.Sleep(time.Duration(1) * time.Second)
	}

	fmt.Println()
	return
}
