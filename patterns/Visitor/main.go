package main

import (
	"visitor/visitor"
)

func main() {
	Square := visitor.Square{Side: 5}
	Triangle := visitor.Triangle{Side1: 2, Side2: 4, Side3: 6}
	Rectangle := visitor.Rectangle{Side1: 3, Side2: 10}

	Perimetr := &visitor.Perimetr{}

	Square.AcceptVisitor(Perimetr)
	Triangle.AcceptVisitor(Perimetr)
	Rectangle.AcceptVisitor(Perimetr)
}
