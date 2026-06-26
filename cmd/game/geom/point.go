package geom

import "math"

// Point is a 2D integer vector.
type Point struct {
	X int
	Y int
}

func P(x, y int) Point {
	return Point{X: x, Y: y}
}

func AllPoint(v int) Point {
	return Point{X: v, Y: v}
}

func (p Point) Elem(index int) int {
	switch index {
	case 0:
		return p.X
	case 1:
		return p.Y
	default:
		return 0
	}
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

func (p Point) Mul(s int) Point {
	return Point{X: p.X * s, Y: p.Y * s}
}

func (p Point) MulFloat(s float64) Vec2 {
	return Vec2{X: float64(p.X) * s, Y: float64(p.Y) * s}
}

func (p Point) MulElem(other Point) Point {
	return Point{X: p.X * other.X, Y: p.Y * other.Y}
}

func (p Point) Div(s int) Point {
	return Point{X: p.X / s, Y: p.Y / s}
}

func (p Point) DivFloat(s float64) Vec2 {
	return Vec2{X: float64(p.X) / s, Y: float64(p.Y) / s}
}

func (p Point) DivElem(other Point) Point {
	return Point{X: p.X / other.X, Y: p.Y / other.Y}
}

func (p Point) Mod(s int) Point {
	return Point{X: p.X % s, Y: p.Y % s}
}

func (p Point) ModElem(other Point) Point {
	return Point{X: p.X % other.X, Y: p.Y % other.Y}
}

func (p Point) Neg() Point {
	return Point{X: -p.X, Y: -p.Y}
}

func (p Point) IsZero() bool {
	return p.X == 0 && p.Y == 0
}

func (p Point) MinComponent() int {
	return min(p.X, p.Y)
}

func (p Point) MaxComponent() int {
	return max(p.X, p.Y)
}

func (p *Point) Clear() {
	p.X = 0
	p.Y = 0
}

func (p Point) WithX(x int) Point {
	return Point{X: x, Y: p.Y}
}

func (p Point) WithY(y int) Point {
	return Point{X: p.X, Y: y}
}

func (p *Point) Set(x, y int) *Point {
	p.X = x
	p.Y = y
	return p
}

func (p Point) With(x, y int) Point {
	p.Set(x, y)
	return p
}

func (p *Point) SetPoint(other Point) *Point {
	*p = other
	return p
}

func (p Point) WithPoint(other Point) Point {
	p.SetPoint(other)
	return p
}

func (p Point) MovedBy(x, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}

func (p Point) MovedByPoint(delta Point) Point {
	return p.Add(delta)
}

func (p *Point) MoveBy(x, y int) *Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) MoveByPoint(delta Point) *Point {
	return p.MoveBy(delta.X, delta.Y)
}

func (p Point) HorizontalAspectRatio() float64 {
	if p.Y == 0 {
		return 0
	}
	return float64(p.X) / float64(p.Y)
}

func (p Point) Length() float64 {
	return math.Sqrt(float64(p.LengthSq()))
}

func (p Point) LengthSq() int {
	return p.X*p.X + p.Y*p.Y
}

func (p Point) ManhattanLength() int {
	return absInt(p.X) + absInt(p.Y)
}

func (p Point) ManhattanDistanceFrom(other Point) int {
	return absInt(p.X-other.X) + absInt(p.Y-other.Y)
}

func (p Point) DistanceFrom(other Point) float64 {
	return math.Sqrt(float64(p.DistanceFromSq(other)))
}

func (p Point) DistanceFromSq(other Point) int {
	x := p.X - other.X
	y := p.Y - other.Y
	return x*x + y*y
}

func (p Point) Area() int {
	return p.X * p.Y
}

func (p Point) Rotated90(n int) Point {
	switch n % 4 {
	case 1, -3:
		return Point{X: -p.Y, Y: p.X}
	case 2, -2:
		return Point{X: -p.X, Y: -p.Y}
	case 3, -1:
		return Point{X: p.Y, Y: -p.X}
	default:
		return p
	}
}

func (p Point) Rotated90At(center Point, n int) Point {
	return p.Sub(center).Rotated90(n).Add(center)
}

func (p *Point) Rotate90(n int) *Point {
	*p = p.Rotated90(n)
	return p
}

func (p *Point) Rotate90At(center Point, n int) *Point {
	*p = p.Rotated90At(center, n)
	return p
}

func (p Point) Angle() float64 {
	if p.IsZero() {
		return math.NaN()
	}
	return math.Atan2(float64(p.X), float64(-p.Y))
}

func (p Point) GetAngle() float64 {
	return p.Angle()
}

func (p Point) Dot(other Point) int {
	return p.X*other.X + p.Y*other.Y
}

func (p Point) Cross(other Point) int {
	return p.X*other.Y - p.Y*other.X
}

func (p Point) AsVec2() Vec2 {
	return Vec2{X: float64(p.X), Y: float64(p.Y)}
}

func (p Point) XX() Point {
	return Point{X: p.X, Y: p.X}
}

func (p Point) XY() Point {
	return p
}

func (p Point) YX() Point {
	return Point{X: p.Y, Y: p.X}
}

func (p Point) YY() Point {
	return Point{X: p.Y, Y: p.Y}
}

func (p Point) X0() Point {
	return Point{X: p.X}
}

func (p Point) Y0() Point {
	return Point{X: p.Y}
}

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
