package structs

type State interface {
	DispenseMoney(money int)
	RequestMoney(money int)
	LoadCash(cash int)
}
