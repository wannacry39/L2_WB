package funcs

import (
	"strconv"
	"strings"
)

func InitArgs(flags []string, del *string, fields []int, sflag *bool) []int {
	for i, val := range flags {
		if val == "-d" {
			*del = flags[i+1]
		}
		if val == "-f" {
			if len(flags[i+1]) > 1 {
				nums := strings.Split(flags[i+1], ",")
				for _, val := range nums {
					intval, _ := strconv.Atoi(val)
					fields = append(fields, intval)
				}
			} else {
				intval, _ := strconv.Atoi(flags[i+1])
				fields = append(fields, intval)
			}
		}
		if val == "-s" {
			*sflag = true
		}

	}
	return fields
}
