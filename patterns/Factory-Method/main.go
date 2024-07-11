package main

import "factory_mthd/structs"

func main() {
	arr := []string{"Gaming", "Office", "Server", "Diff_type"}

	for _, val := range arr {
		pc := structs.NewPC(val)

		if pc != nil {
			pc.Configuration()
		}
	}
}
