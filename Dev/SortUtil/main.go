package main

import (
	"errors"
	"fmt"
	"funcs/funcs"
	"os"
	"strconv"
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v\n", errors.New("openning file failed"))
		os.Exit(1)
	}
	defer filePtr.Close()

	funcs.SortFile(filePtr, table, kvalue)

}
