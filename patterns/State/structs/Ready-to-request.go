package structs

import "fmt"

type ReadyForRequest struct {
	atm *ATM
}

func (r *ReadyForRequest) DispenseMoney(money int) {
	fmt.Println("Please set an amount of cash")
	return
}

func (r *ReadyForRequest) RequestMoney(money int) {
	if money <= r.atm.Cash {
		fmt.Println("Success! Dispensing money...")
		r.atm.SetState(r.atm.MoneyDispensing)
		return
	}
	fmt.Println("Sorry, ATM has not enough cash...")
	r.atm.SetState(r.atm.NotEnoughCash)
	return

}

func (r *ReadyForRequest) LoadCash(cash int) {
	r.atm.Cash += cash
	fmt.Printf("Loaded %d\n", cash)
	return
}
