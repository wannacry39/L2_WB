package main

import (
	"states/structs"
	"time"
)

func main() {
	ATM := structs.NewATM(2000)
	money := 1500
	ATM.RequestMoney(money)
	time.Sleep(3 * time.Second)
	ATM.DispenseMoney(money)
	ATM.RequestMoney(money)
	ATM.LoadCash(1000)
	time.Sleep(3 * time.Second)
	ATM.RequestMoney(money)
	time.Sleep(3 * time.Second)
	ATM.DispenseMoney(money)

}
