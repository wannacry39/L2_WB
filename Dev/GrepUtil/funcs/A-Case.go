package funcs

import (
	"strconv"
	"strings"
)

func A_case(data []string, templ string, table map[string]struct{}, res []string, count int, d int) []string {
	for i, val := range data {
		if _, ok := table["-F"]; ok {
			if val == templ {
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				count++
				if len(data) < d {
					d = len(data) - i
				} else {
					d++
				}
				res = append(res, val)
				res = append(res, data[i+1:i+d]...)
				continue
			}
		}
		if _, ok := table["-i"]; ok {
			if strings.Contains(strings.ToLower(val), strings.ToLower(templ)) {
				count++
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				if len(data)-i < d {
					d = len(data) - i
				} else {
					d++
				}
				res = append(res, val)
				res = append(res, data[i+1:i+d]...)
				continue
			}
		}
		if strings.Contains(val, templ) {
			count++
			if _, ok := table["-n"]; ok {
				val = strconv.Itoa(i+1) + ")" + val
			}
			if len(data)-i < d {
				d = len(data) - i
			} else {
				d++
			}
			res = append(res, val)
			res = append(res, data[i+1:i+d]...)

		}
	}
	return res
}
