package iterations

import (
	"strings"
)

func Repeat(char string, times int) string {
	var res strings.Builder

	for i := 0; i < times; i++ {
		res.WriteString(char)
	}
	return res.String()
}
