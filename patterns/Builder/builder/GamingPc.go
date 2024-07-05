package builder

type GamingPcBuilder struct {
	CPU     string
	MEM     int
	Gr_Card string
	brand   string
}

func (g *GamingPcBuilder) SetCPU() {
	g.CPU = "Intel Core I7 6700K"
}

func (g *GamingPcBuilder) SetMEM() {
	g.MEM = 16
}

func (g *GamingPcBuilder) SetCard() {
	g.Gr_Card = "Nvidia RTX 4090"
}

func (g *GamingPcBuilder) SetBrand() {
	g.brand = "Ultra Mega PC"
}

func (g *GamingPcBuilder) GetPC() PC {
	return PC{
		CPU:     g.CPU,
		Mem:     g.MEM,
		Gr_Card: g.Gr_Card,
		brand:   g.brand,
	}
}
