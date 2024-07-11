package structs

import "fmt"

type NotEnoughCash struct {
	atm *ATM
}

func (n *NotEnoughCash) DispenseMoney(money int) {
	fmt.Println("Sorry, ATM has not enough cash...")
}

func (n *NotEnoughCash) RequestMoney(money int) {
	if money <= n.atm.Cash {
		fmt.Println("Success! Dispensing money...")
		n.atm.SetState(n.atm.MoneyDispensing)
		return
	}
	fmt.Println("Sorry, ATM has not enough cash...")

}

func (n *NotEnoughCash) LoadCash(cash int) {
	n.atm.Cash += cash
	fmt.Printf("Loaded %d\n", cash)
	n.atm.SetState(n.atm.ReadyForRequest)
	return
}
