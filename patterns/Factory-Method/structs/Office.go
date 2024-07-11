package structs

import "fmt"

type OfficePC struct {
	Type     string
	CPU      int
	MEM      int
	Gr_Card  string
	Monitors int
}

func NewOfficePC() PC {
	return OfficePC{
		Type:     "OfficePC",
		CPU:      4,
		MEM:      4,
		Gr_Card:  "None",
		Monitors: 1,
	}
}

func (o OfficePC) Configuration() {
	fmt.Printf("Type: %s, CPU: %d, MEM: %d, GRAPHICS: %s, Monitors: %d\n", o.Type, o.CPU, o.MEM, o.Gr_Card, o.Monitors)
}
