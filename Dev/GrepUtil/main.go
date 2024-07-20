package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type GrepFlags struct {
	flagsMap map[string]struct{}
}

var flagsTable = GrepFlags{flagsMap: map[string]struct{}{
	"-A": {},
	"-B": {},
	"-C": {},
	"-c": {},
	"-i": {},
	"-v": {},
	"-F": {},
	"-n": {},
}}

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
		ScanFunc(os.Stdin, flags, names[0], table)
		return
	}

	file, err := os.Open(names[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in openning file: %v\n", err)
		os.Exit(1)
	}

	ScanFunc(file, flags, names[0], table)
}

func ScanFunc(r io.Reader, flags []string, templ string, table map[string]struct{}) {
	scn := bufio.NewScanner(r)
	data := []string{}
	res := []string{}
	for scn.Scan() {
		data = append(data, scn.Text()+"\n")
	}
	for i := 0; i < len(flags); i++ {
		if d, err := strconv.Atoi(flags[i]); err == nil {
			switch flags[i-1] {
			case "-A":
				for i, val := range data {
					if strings.Contains(val, templ) {
						if len(data)-i < d {
							d = len(data) - i
						}
						res = append(res, val)
						res = append(res, data[i+1:i+1+d]...)

					}
				}
			case "-B":
				for i, val := range data {
					if strings.Contains(val, templ) {
						if i < d {
							d = i
						}
						res = append(res, data[i-d:i]...)
						res = append(res, val)

					}
				}
			case "-C":
				for i, val := range data {
					if strings.Contains(val, templ) {
						res = append(res, data[i-d:i]...)
						res = append(res, val)
						res = append(res, data[i+1:i+d+1]...)

					}
				}
			}
		}
	}
	fmt.Print(res)
}
