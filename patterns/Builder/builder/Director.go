package builder

type Director struct {
	Builder Builder
}

func NewDir(builder Builder) *Director {
	return &Director{Builder: builder}
}

func (d *Director) SetBuilder(builder Builder) {
	d.Builder = builder
}

func (d *Director) CreatePC() PC {
	d.Builder.SetCPU()
	d.Builder.SetMEM()
	d.Builder.SetCard()
	d.Builder.SetBrand()
	return d.Builder.GetPC()
}
