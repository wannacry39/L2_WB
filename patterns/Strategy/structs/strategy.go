package structs

type Strategy interface {
	Route(start, end float64)
}
