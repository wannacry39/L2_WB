package structs

import "fmt"

type ReadyForDispense struct {
	atm *ATM
}

func (m *ReadyForDispense) DispenseMoney(money int) {
	m.atm.Cash -= money
	fmt.Printf("Take your %d\n", money)
	m.atm.SetState(m.atm.ReadyForRequest)
}

func (m *ReadyForDispense) RequestMoney(money int) {
	fmt.Println("Wait till the end of the operation")
	return
}

func (m *ReadyForDispense) LoadCash(cash int) {
	fmt.Println("Wait till the end of the operation")
	return
}
