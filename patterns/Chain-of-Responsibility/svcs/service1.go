package svcs

import "fmt"

type GetData struct {
	Name string
	Next Service
}

func (g *GetData) Execute(data *Data) {
	if data.GotData {
		fmt.Println("Got Data, sending to the next service...")
		g.Next.Execute(data)
		return
	}
	fmt.Println("Got Data, setting flag and then sending to the next service...")
	data.GotData = true
	g.Next.Execute(data)

}

func (g *GetData) SetNext(svc Service) {
	g.Next = svc
}
