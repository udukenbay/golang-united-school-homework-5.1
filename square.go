package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sq Square) End() Point {
	x := sq.start.x + int(sq.a)
	y := sq.start.y + int(sq.a)
	return Point{x, y}
}

func (sq Square) Area() uint {
	return sq.a * sq.a
}

func (sq Square) Perimeter() uint {
	return 4 * sq.a
}
