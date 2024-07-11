package main

import "structs/structs"

func main() {
	arr := []structs.Strategy{
		structs.PublicTransportStrategy{},
		structs.RoadStrategy{},
		structs.WalkStrategy{}}

	nav := structs.Navigator{}

	for _, str := range arr {
		nav.SetStrategy(str)
		nav.Route(8.56, 16.56)
	}
}
