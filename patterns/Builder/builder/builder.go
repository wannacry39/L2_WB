package builder

type Builder interface {
	SetCPU()
	SetMEM()
	SetCard()
	SetBrand()
	GetPC() PC
}

func GetBuilder(_type string) Builder {
	switch _type {
	default:
		return nil
	case "GamingPC":
		return &GamingPcBuilder{}
	case "OfficePC":
		return &OfficePcBuilder{}
	}
}
