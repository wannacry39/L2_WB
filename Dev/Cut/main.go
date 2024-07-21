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
	fields := 0
	for i, val := range flags {
		if val == "-d" {
			delimiter = flags[i+1]
		}
		if val == "-f" {
			fields, _ = strconv.Atoi(flags[i+1])
		}
	}
	scn := bufio.NewScanner(os.Stdin)
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
	fmt.Println(data)
	if fields != 0 {
		if v, ok := data[fields]; ok {
			for _, word := range v {
				fmt.Println(word)
			}
		}
	}
}
