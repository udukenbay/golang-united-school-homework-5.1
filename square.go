package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sq *Square) End() Point {
	var p Point = Point{x: 1, y: 2}
	p.x = p.x + int(sq.a)
	p.y = p.y - int(sq.a)
	return p
}

func (sq *Square) Area() uint {
	return sq.a * sq.a
}

func (sq *Square) Perimeter() uint {
	return 4 * sq.a
}
