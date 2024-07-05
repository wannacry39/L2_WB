package builder

type OfficePcBuilder struct {
	CPU     string
	MEM     int
	Gr_Card string
	brand   string
}

func (g *OfficePcBuilder) SetCPU() {
	g.CPU = "Intel Core I3 5300"
}

func (g *OfficePcBuilder) SetMEM() {
	g.MEM = 4
}

func (g *OfficePcBuilder) SetCard() {
	g.Gr_Card = "Intel HD Graphics"
}

func (g *OfficePcBuilder) SetBrand() {
	g.brand = "Just Office PC"
}

func (g *OfficePcBuilder) GetPC() PC {
	return PC{
		CPU:     g.CPU,
		Mem:     g.MEM,
		Gr_Card: g.Gr_Card,
		brand:   g.brand,
	}
}
