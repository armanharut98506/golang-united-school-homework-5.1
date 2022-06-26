package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	end := Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
	return end
}

func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return 4 * s.a
}me
