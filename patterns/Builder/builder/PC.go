package builder

import "fmt"

type PC struct {
	CPU     string
	Mem     int
	Gr_Card string
	brand   string
}

func (pc PC) PrintConfig() {
	fmt.Printf("PC configuration:\nBRAND: %s\nCPU: %s\nMEM: %d\nGRAPHIC: %s\n", pc.brand, pc.CPU, pc.Mem, pc.Gr_Card)
}
