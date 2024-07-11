package structs

import "fmt"

type ServerPC struct {
	Type     string
	CPU      int
	MEM      int
	Gr_Card  string
	Monitors int
}

func NewServerPC() PC {
	return ServerPC{
		Type:     "GamingPC",
		CPU:      16,
		MEM:      64,
		Gr_Card:  "Ultra server graphic card",
		Monitors: 0,
	}
}

func (s ServerPC) Configuration() {
	fmt.Printf("Type: %s, CPU: %d, MEM: %d, GRAPHICS: %s, Monitors: %d\n", s.Type, s.CPU, s.MEM, s.Gr_Card, s.Monitors)
}
