package funcs

import (
	"strconv"
	"strings"
)

func B_case(data []string, templ string, table map[string]struct{}, res []string, count int, d int) []string {
	for i, val := range data {
		if _, ok := table["-F"]; ok {
			if val == templ {
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				count++
				if i < d {
					d = i
				}
				res = append(res, data[i-d:i]...)
				res = append(res, val)
				continue
			}
		}
		if _, ok := table["-i"]; ok {
			if strings.Contains(strings.ToLower(val), strings.ToLower(templ)) {
				count++
				if _, ok := table["-n"]; ok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				if i < d {
					d = i
				}
				res = append(res, data[i-d:i]...)
				res = append(res, val)
				continue
			}
		}

		if strings.Contains(val, templ) {
			count++
			if _, ok := table["-n"]; ok {
				val = strconv.Itoa(i+1) + ")" + val
			}
			if i < d {
				d = i
			}
			res = append(res, data[i-d:i]...)
			res = append(res, val)

		}
	}
	return res
}
