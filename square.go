package square

//package main

import "fmt"

func measure(g geometry) {
	fmt.Println(g)
	fmt.Printf("area     : %f \n", g.Area())
	fmt.Printf("perimeter: %f \n", g.Perimeter())
}

func main() {
	p := Point{3, 5}
	s := Square{a: 4}
	fmt.Printf("%.3f\n", s.End(p))
	measure(&s)
}

type geometry interface {
	Area() uint
	Perimeter() uint
}

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sq Square) End(p Point) Point {
	p.x = p.x + int(sq.a)
	p.y = p.y - int(sq.a)
	return p
}

func (sq *Square) Area() uint {
	return sq.a * sq.a
}

func (sq Square) Perimeter() uint {
	return 4 * sq.a
}
