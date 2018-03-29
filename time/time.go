package time

import "fmt"

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
