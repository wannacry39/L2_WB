package structs

import "fmt"

type RoadStrategy struct {
}

func (r RoadStrategy) Route(start, end float64) {
	AvgSpeed := 45
	TrafficFactor := 5
	d := end - start
	time := d / float64(AvgSpeed) * float64(TrafficFactor)
	fmt.Printf("Distance from START[%f] to the END[%f] is %f. Estimated time of arrival: %f\n", start, end, d, time)

}
