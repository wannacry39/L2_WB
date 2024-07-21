package funcs

import (
	"strconv"
	"strings"
)

func MainCase(data []string, templ string, table map[string]struct{}, res []string, count int) []string {
	for i, val := range data {
		if _, ok := table["-F"]; ok {
			if val[:len(val)-1] == templ {
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				count++
				res = append(res, val)

			}
			continue
		}
		if _, ok := table["-i"]; ok {
			if strings.Contains(strings.ToLower(val), strings.ToLower(templ)) {
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				count++
				res = append(res, val)

			}
			continue
		}
		if strings.Contains(val, templ) {
			if _, ok := table["-n"]; ok {
				val = strconv.Itoa(i+1) + ")" + val
			}
			count++
			res = append(res, val)
		}
	}
	return res
}
