package funcs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SortFile(f *os.File, t map[string]struct{}, kval int) {
	scn := bufio.NewScanner(f)
	if _, ok := t["-n"]; ok {
		if _, ok := t["-k"]; ok {
			data := make(map[int][]int)
			for scn.Scan() {
				line := strings.Fields(scn.Text())
				for i, val := range line {
					chislo, err := strconv.Atoi(val)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error:%v\n", errors.New("Please enter file with int values."))
						os.Exit(1)
					}
					data[i] = append(data[i], chislo)
				}
			}
			URflags(data, t, kval)

		}
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

		slices.Sort(ints)

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
		URflags(data, t, kval)
	}
	strs := []string{}
	for scn.Scan() {
		strs = append(strs, scn.Text())
	}

	if _, ok := t["-u"]; ok {
		strs = isUnique(strs)
	}

	slices.Sort(strs)

	if _, ok := t["-r"]; ok {
		Reverse(strs)
	}
	for _, val := range strs {
		fmt.Println(val)
	}
	return
}
