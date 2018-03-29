package time

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// TimeHms は秒数値をH,M,Sの数値に変換します。
func TimeHms(i int) (int, int, int) {
	h := i / 60 / 60
	m := i / 60 % 60
	s := i % 60
	return h, m, s
}

// TimeHmsString は秒単位の数値をHH:MM:SSの文字列に変換します。
func TimeHmsString(i int) string {
	h, m, s := TimeHms(i)
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// TimeToSec はHH:MM:SSを秒に変換します。
func TimeSec(t string) (sec int, err error) {
	hms := strings.Split(t, ":")

	h, err := strconv.Atoi(hms[0])
	if err != nil {
		log.Println(err)
		return
	}
	sec += h * 60 * 60

	m, err := strconv.Atoi(hms[1])
	if err != nil {
		log.Println(err)
		return
	}
	sec += m * 60

	s, err := strconv.Atoi(hms[2])
	if err != nil {
		log.Println(err)
		return
	}
	sec += s

	return
}
