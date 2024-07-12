package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	TimeNow, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	time := time.Now()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v / %v\n", time, TimeNow)
}
