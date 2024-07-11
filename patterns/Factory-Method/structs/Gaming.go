package structs

import "fmt"

type GamingPC struct {
	Type     string
	CPU      int
	MEM      int
	Gr_Card  string
	Monitors int
}

func NewGamingPC() PC {
	return GamingPC{
		Type:     "GamingPC",
		CPU:      8,
		MEM:      16,
		Gr_Card:  "Nvidia RTX 4090",
		Monitors: 2,
	}
}

func (g GamingPC) Configuration() {
	fmt.Printf("Type: %s, CPU: %d, MEM: %d, GRAPHICS: %s, Monitors: %d\n", g.Type, g.CPU, g.MEM, g.Gr_Card, g.Monitors)
}
