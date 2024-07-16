package funcs

import (
	"fmt"
	"slices"
)

func Reverse[V int | string](arr []V) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

func isUnique[V int | string](arr []V) []V {
	tmp := make(map[V]struct{})
	res := []V{}
	for _, val := range arr {
		if _, ok := tmp[val]; ok {
			continue
		}
		res = append(res, val)
		tmp[val] = struct{}{}
	}
	return res
}

func Max[V int | string](t map[int][]V) int {
	res := 0
	for _, v := range t {
		if res < len(v) {
			res = len(v)
		}
	}
	return res
}

func URflags[V int | string](data map[int][]V, t map[string]struct{}, kval int) {
	if _, ok := t["-u"]; ok {
		data[kval] = isUnique(data[kval])
	}

	slices.Sort(data[kval])

	if _, ok := t["-r"]; ok {
		Reverse(data[kval])
	}

	for j := 0; j < Max(data); j++ {

		for i := 0; i < len(data); i++ {
			if j >= len(data[i]) {
				continue
			}
			fmt.Printf("%v ", data[i][j])
		}
		fmt.Print("\n")
	}
	return
}
