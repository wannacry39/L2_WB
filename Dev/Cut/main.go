package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flags := os.Args[1:]
	delimiter := "\t"
	fields := []int{}
	sflag := false
	for i, val := range flags {
		if val == "-d" {
			delimiter = flags[i+1]
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
			sflag = true
		}

	}

	scn := bufio.NewScanner(os.Stdin)
	if sflag {
		lines := []string{}
		for scn.Scan() {
			line := scn.Text()
			if line == "" {
				break
			}
			lines = append(lines, line)
		}
		for _, val := range lines {
			if strings.Contains(val, delimiter) {
				fmt.Println(val)
			}
		}
		return
	}
	data := make(map[int][]string)
	for scn.Scan() {
		line := scn.Text()
		if line == "" {
			break
		}
		tmp := strings.Split(line, delimiter)
		for i, val := range tmp {
			if _, ok := data[i+1]; ok {
				data[i+1] = append(data[i+1], val)
				continue
			}
			data[i+1] = []string{val}
		}
	}

	for _, f := range fields {
		if v, ok := data[f]; ok {
			fmt.Printf("Column: %d\n", f)
			for _, word := range v {

				fmt.Printf("%s ", word)
				fmt.Println()
			}
		}
	}
}
