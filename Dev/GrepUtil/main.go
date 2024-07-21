package main

import (
	"Greputil/funcs"
	"fmt"
	"os"
	"strconv"
)

func main() {

	flags := []string{}
	names := []string{}
	table := make(map[string]struct{})
	for i := 1; i < len(os.Args); i++ {
		if _, ok := flagsTable.flagsMap[os.Args[i]]; ok {
			flags = append(flags, os.Args[i])
			table[os.Args[i]] = struct{}{}
			if _, err := strconv.Atoi(os.Args[i+1]); err == nil {
				flags = append(flags, os.Args[i+1])
				i++
			}
			continue
		}
		if os.Args[i][0] != '-' {
			names = append(names, os.Args[i])
			continue
		}
		fmt.Fprint(os.Stderr, "Check usage. Invalid args\n")
		os.Exit(1)
	}

	if len(names) == 1 {
		funcs.ScanFunc(os.Stdin, flags, names[0], table)
		return
	}

	file, err := os.Open(names[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in openning file: %v\n", err)
		os.Exit(1)
	}

	funcs.ScanFunc(file, flags, names[0], table)
}
