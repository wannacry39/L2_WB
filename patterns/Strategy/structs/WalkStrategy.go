package structs

import "fmt"

type WalkStrategy struct {
}

func (w WalkStrategy) Route(start, end float64) {
	AvgSpeed := 3
	d := end - start
	time := d / float64(AvgSpeed)
	fmt.Printf("Distance from START[%f] to the END[%f] is %f. Estimated time of arrival: %f\n", start, end, d, time)

}
