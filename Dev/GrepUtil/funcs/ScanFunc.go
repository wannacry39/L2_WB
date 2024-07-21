package funcs

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

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
	if !Aok && !Bok && !Cok {
		res = MainCase(data, templ, table, res, count)
	}

	for i := 0; i < len(flags); i++ {
		if d, err := strconv.Atoi(flags[i]); err == nil {
			switch flags[i-1] {
			case "-A":
				res = A_case(data, templ, table, res, count, d)
			case "-B":
				res = B_case(data, templ, table, res, count, d)
			case "-C":
				res = C_case(data, templ, table, res, count, d)
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
