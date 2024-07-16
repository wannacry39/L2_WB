package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	dict := []string{"тяпка", "пятак", "пятка", "листок", "столик", "слиток", "Привет", "Вор", "Ров"}

	Anagrams := *AnagramSearch(&dict)
	for k, v := range Anagrams {
		fmt.Printf("Key: %s Value: ", k)
		for _, word := range *v {
			fmt.Printf("%s ", word)
		}
		fmt.Print("\n")
	}
}

func AnagramSearch(arr *[]string) *map[string]*[]string {
	res := make(map[string]*[]string)
	check_table := make(map[string]string)

	for _, val := range *arr {
		val := strings.ToLower(val)
		tmp := strings.Split(val, "")
		sort.Strings(tmp)
		sorted := strings.Join(tmp, "")
		if v, ok := check_table[sorted]; ok {
			*res[v] = append(*res[v], val)
			sort.Strings(*res[v])
			continue
		}
		check_table[sorted] = val
		res[val] = &[]string{}
	}
	return &res
}
