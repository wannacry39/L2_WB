package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	table := make(map[string]struct{})
	kvalue := 0
	for i, val := range args {
		table[val] = struct{}{}
		if val == "-k" {
			kvalue, _ = strconv.Atoi(args[i+1])
		}
	}

	filePtr, err := os.Open(args[len(args)-1])
	defer filePtr.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v", errors.New("Some error in opening file."))
		os.Exit(1)
	}

	SortFile(filePtr, table, kvalue)

}

func SortFile(f *os.File, t map[string]struct{}, kval int) {
	scn := bufio.NewScanner(f)
	if _, ok := t["-n"]; ok {
		ints := []int{}
		for scn.Scan() {
			line, err := strconv.Atoi(scn.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error:%v\n", errors.New("Please enter file with int values."))
				os.Exit(1)
			}
			ints = append(ints, line)
		}
		if _, ok := t["-u"]; ok {
			ints = isUnique(ints)
		}

		sort.Ints(ints)

		if _, ok := t["-r"]; ok {
			Reverse(ints)
		}

		for _, val := range ints {
			fmt.Println(val)
		}
		return
	}
	if _, ok := t["-k"]; ok {
		data := make(map[int][]string)
		for scn.Scan() {
			line := strings.Fields(scn.Text())
			for i, val := range line {
				data[i] = append(data[i], val)
			}
		}
		if _, ok := t["-u"]; ok {
			data[kval] = isUnique(data[kval])
		}

		sort.Strings(data[kval])

		if _, ok := t["-r"]; ok {
			Reverse(data[kval])
		}

		for j := 0; j < Max(data); j++ {

			for i := 0; i < len(data); i++ {
				if j >= len(data[i]) {
					continue
				}
				fmt.Printf("%s ", data[i][j])
			}
			fmt.Print("\n")
		}

		return

	}
	strs := []string{}
	for scn.Scan() {
		strs = append(strs, scn.Text())
	}

	if _, ok := t["-u"]; ok {
		strs = isUnique(strs)
	}

	sort.Strings(strs)

	if _, ok := t["-r"]; ok {
		Reverse(strs)
	}
	for _, val := range strs {
		fmt.Println(val)
	}
	return
}

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

func Max(t map[int][]string) int {
	res := 0
	for _, v := range t {
		if res < len(v) {
			res = len(v)
		}
	}
	return res
}
