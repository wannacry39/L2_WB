package structs

import "fmt"

type PC interface {
	Configuration()
}

func NewPC(typename string) PC {
	switch typename {
	case "Server":
		return NewServerPC()
	case "Gaming":
		return NewGamingPC()
	case "Office":
		return NewOfficePC()
	default:
		fmt.Println("Unknown type")
		return nil

	}
}
