package svcs

import "fmt"

type Saving struct {
	Name string
	Next Service
}

func (s *Saving) Execute(data *Data) {
	if !data.UpdateData {
		fmt.Println("Data is not updated")
		return
	}
	fmt.Println("Saving data...")

}

func (upd *Saving) SetNext(svc Service) {
	upd.Next = svc
}
