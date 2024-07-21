package funcs

import (
	"strconv"
	"strings"
)

func C_case(data []string, templ string, table map[string]struct{}, res []string, count int, d int) []string {
	dpos := d
	for i, val := range data {
		if _, ok := table["-F"]; ok {
			if val == templ {
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				count++
				if d > i {
					d = i
				}
				if dpos > len(data)-i {
					dpos = len(data) - i
				} else {
					dpos++
				}
				res = append(res, data[i-d:i]...)
				res = append(res, val)
				res = append(res, data[i+1:i+dpos]...)
				continue
			}
		}
		if _, ok := table["-i"]; ok {
			if strings.Contains(strings.ToLower(val), strings.ToLower(templ)) {
				count++
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				if d > i {
					d = i
				}
				if dpos > len(data)-i {
					dpos = len(data) - i
				} else {
					dpos++
				}
				res = append(res, data[i-d:i]...)
				res = append(res, val)
				res = append(res, data[i+1:i+dpos]...)
				continue
			}
		}
		if strings.Contains(val, templ) {
			count++
			if _, ok := table["-n"]; ok {
				val = strconv.Itoa(i+1) + ")" + val
			}
			if d > i {
				d = i
			}
			if dpos >= len(data)-i {
				dpos = len(data) - i
			} else {
				dpos++
			}
			res = append(res, data[i-d:i]...)
			res = append(res, val)
			res = append(res, data[i+1:i+dpos]...)

		}
	}
	return res
}
