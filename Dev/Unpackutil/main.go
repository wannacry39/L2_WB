package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	res, err := Unpack(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res)
}

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", errors.New("invalid string")
	}
	res := make([]string, len(s))
	arr_s := []rune(s)
	for i := 0; i < len(arr_s); i++ {
		if string(arr_s[i]) == "\\" {
			res = append(res, string(arr_s[i+1]))
			i++
			continue
		}
		if _, err := strconv.Atoi(string(arr_s[i])); err == nil {
			chislo := []byte{}
			chislo = append(chislo, byte(arr_s[i]))
			j := i + 1
			for j < len(arr_s) {
				if _, err := strconv.Atoi(string(arr_s[j])); err != nil {
					break
				}
				chislo = append(chislo, byte(arr_s[j]))
				j++
			}
			int_chislo, _ := strconv.Atoi(string(chislo))
			repeat := make([]string, int_chislo)
			for int_chislo-1 > 0 {
				repeat = append(repeat, string(arr_s[i-1]))
				int_chislo--
			}
			res = append(res, repeat...)
			continue
		}
		res = append(res, string(arr_s[i]))
	}
	return strings.Join(res, ""), nil
}
