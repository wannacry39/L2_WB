package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	filePtr, err := os.Open(args[len(args)-1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v", errors.New("Some error in opening file."))
		os.Exit(1)
	}
	SortFile(filePtr, args)
	defer filePtr.Close()
}

func SortFile(f *os.File, s []string) {
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		for _, key := range s {
			switch key {
			case "-k":
				arr := strings.Fields(scn.Text())
			case "-n":

			case "-r":

			case "-u":

			default:
				continue
			}
		}
	}

}
