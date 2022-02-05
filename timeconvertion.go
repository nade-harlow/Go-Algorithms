package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Timeconvertion("02:34:50PM"))
}

func Timeconvertion(s string) string {

	hourStr := string(s[0]) + string(s[1])
	minuteStr := string(s[3]) + string(s[4])
	secondsStr := string(s[6]) + string(s[7])
	hour, _ := strconv.Atoi(hourStr)
	if strings.HasSuffix(s, "AM") && hour == 12 {
		hour = 0
	}
	if strings.HasSuffix(s, "PM") && hour != 12 {
		hour += 12
	}
	return fmt.Sprintf("%02d:%s:%s", hour, minuteStr, secondsStr)
}
