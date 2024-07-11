package svcs

import "fmt"

type UpdateData struct {
	Name string
	Next Service
}

func (upd *UpdateData) Execute(data *Data) {
	if data.UpdateData {
		fmt.Println("Data is already updated, sending to the next service...")
		upd.Next.Execute(data)
		return
	}
	fmt.Println("Updating data, and sending to the next service...")
	data.UpdateData = true
	upd.Next.Execute(data)

}

func (upd *UpdateData) SetNext(svc Service) {
	upd.Next = svc
}
