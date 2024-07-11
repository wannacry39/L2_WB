package structs

type ATM struct {
	NotEnoughCash   State
	ReadyForRequest State
	MoneyDispensing State
	CurrState       State
	Cash            int
}

func NewATM(cash int) *ATM {
	atm := &ATM{Cash: cash}

	NotEnough := &NotEnoughCash{
		atm: atm,
	}
	ReadyForRqst := &ReadyForRequest{
		atm: atm,
	}

	ReadyForDspns := &ReadyForDispense{
		atm: atm,
	}

	atm.SetState(ReadyForRqst)
	atm.ReadyForRequest = ReadyForRqst
	atm.NotEnoughCash = NotEnough
	atm.MoneyDispensing = ReadyForDspns
	return atm
}

func (a *ATM) DispenseMoney(money int) {
	a.CurrState.DispenseMoney(money)
}

func (a *ATM) RequestMoney(money int) {
	a.CurrState.RequestMoney(money)
}

func (a *ATM) SetState(s State) {
	a.CurrState = s
}

func (a *ATM) LoadCash(cash int) {
	a.CurrState.LoadCash(cash)
}
