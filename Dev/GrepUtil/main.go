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
	count := 0
	for scn.Scan() {
		data = append(data, scn.Text()+"\n")
	}

	_, Aok := table["-A"]
	_, Bok := table["-B"]
	_, Cok := table["-C"]
	_, nok := table["-n"]
	if !Aok && !Bok && !Cok {
		for i, val := range data {
			if _, ok := table["-F"]; ok {
				if val[:len(val)-1] == templ {
					if nok {
						val = strconv.Itoa(i+1) + ")" + val
					}
					count++
					res = append(res, val)

				}
				continue
			}
			if _, ok := table["-i"]; ok {
				if strings.Contains(strings.ToLower(val), strings.ToLower(templ)) {
					if nok {
						val = strconv.Itoa(i+1) + ")" + val
					}
					count++
					res = append(res, val)

				}
				continue
			}
			if strings.Contains(val, templ) {
				if nok {
					val = strconv.Itoa(i+1) + ")" + val
				}
				count++
				res = append(res, val)
			}
		}
	}

	for i := 0; i < len(flags); i++ {
		if d, err := strconv.Atoi(flags[i]); err == nil {
			switch flags[i-1] {
			case "-A":
				for i, val := range data {
					if _, ok := table["-F"]; ok {
						if val == templ {
							if nok {
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
							if nok {
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
						if nok {
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
			case "-B":
				for i, val := range data {
					if _, ok := table["-F"]; ok {
						if val == templ {
							if nok {
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
							if nok {
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
						if nok {
							val = strconv.Itoa(i+1) + ")" + val
						}
						if i < d {
							d = i
						}
						res = append(res, data[i-d:i]...)
						res = append(res, val)

					}
				}
			case "-C":
				dpos := d
				for i, val := range data {
					if _, ok := table["-F"]; ok {
						if val == templ {
							if nok {
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
							if nok {
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
						if nok {
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
			}
		}

	}
	if _, ok := table["-c"]; ok {
		defer fmt.Printf("Rows found: %d\n", count)
	}
	if _, ok := table["-v"]; ok {
		bigString := strings.Join(res, "")
		for _, val := range data {
			if strings.Contains(bigString, val) {
				continue
			}
			fmt.Print(val)
		}
		fmt.Println("Mod: Inverted")
		return
	}
	for _, val := range res {
		fmt.Print(val)
	}

}
