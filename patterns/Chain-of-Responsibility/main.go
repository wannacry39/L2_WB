package main

import "chain/svcs"

func main() {
	svc1 := svcs.GetData{Name: "Service_1"}
	svc2 := svcs.UpdateData{Name: "Service_2"}
	svc3 := svcs.Saving{Name: "Service_3"}

	data := svcs.Data{Id: "Your data"}
	svc1.SetNext(&svc2)
	svc2.SetNext(&svc3)

	svc1.Execute(&data)
}
