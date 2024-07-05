package visitor

import "fmt"

type Visitor interface {
	VisitForSquare(Square)
	VisitForTriangle(Triangle)
	VisitForRectangle(Rectangle)
}

type Perimetr struct {
	val int
}

func (p *Perimetr) VisitForSquare(s Square) {
	p.val = s.Side * 4
	fmt.Printf("Perimetr of Square: %d\n", p.val)
}

func (p *Perimetr) VisitForTriangle(t Triangle) {
	p.val = t.Side1 + t.Side2 + t.Side3
	fmt.Printf("Perimetr of Triangle: %d\n", p.val)
}

func (p *Perimetr) VisitForRectangle(r Rectangle) {
	p.val = (r.Side1 + r.Side2) * 2
	fmt.Printf("Perimetr of Rectangle: %d\n", p.val)
}

type Shape interface {
	GetType()
	AcceptVisitor(Visitor)
}

type Square struct {
	Side int
}

func (s Square) GetType() {
	fmt.Println("Квадрат")
}

func (s Square) AcceptVisitor(v Visitor) {
	v.VisitForSquare(s)
}

type Triangle struct {
	Side1 int
	Side2 int
	Side3 int
}

func (t Triangle) GetType() {
	fmt.Println("Треугольник")
}

func (t Triangle) AcceptVisitor(v Visitor) {
	v.VisitForTriangle(t)
}

type Rectangle struct {
	Side1 int
	Side2 int
}

func (r Rectangle) GetType() {
	fmt.Println("Прямоугольник")
}

func (r Rectangle) AcceptVisitor(v Visitor) {
	v.VisitForRectangle(r)
}
