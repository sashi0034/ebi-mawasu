package geom

// Rect is a rectangle with integer components.
type Rect struct {
	X int
	Y int
	W int
	H int
}

func NewRect(x, y, w, h int) Rect {
	return Rect{X: x, Y: y, W: w, H: h}
}

func EmptyRect() Rect {
	return Rect{}
}

func RectFromPoints(a, b Point) Rect {
	x0, x1 := min(a.X, b.X), max(a.X, b.X)
	y0, y1 := min(a.Y, b.Y), max(a.Y, b.Y)
	return Rect{X: x0, Y: y0, W: x1 - x0, H: y1 - y0}
}

func (r Rect) Pos() Point {
	return Point{X: r.X, Y: r.Y}
}

func (r Rect) Size() Point {
	return Point{X: r.W, Y: r.H}
}

func (r Rect) IsEmpty() bool {
	return r.W == 0 || r.H == 0
}

func (r Rect) HasArea() bool {
	return r.W != 0 && r.H != 0
}

func (r Rect) LeftX() int {
	return r.X
}

func (r Rect) RightX() int {
	return r.X + r.W
}

func (r Rect) TopY() int {
	return r.Y
}

func (r Rect) BottomY() int {
	return r.Y + r.H
}

func (r Rect) CenterX() float64 {
	return float64(r.X) + float64(r.W)*0.5
}

func (r Rect) CenterY() float64 {
	return float64(r.Y) + float64(r.H)*0.5
}

func (r Rect) TL() Point {
	return r.Pos()
}

func (r Rect) TR() Point {
	return Point{X: r.RightX(), Y: r.Y}
}

func (r Rect) BR() Point {
	return Point{X: r.RightX(), Y: r.BottomY()}
}

func (r Rect) BL() Point {
	return Point{X: r.X, Y: r.BottomY()}
}

func (r Rect) TopCenter() Vec2 {
	return Vec2{X: r.CenterX(), Y: float64(r.Y)}
}

func (r Rect) RightCenter() Vec2 {
	return Vec2{X: float64(r.RightX()), Y: r.CenterY()}
}

func (r Rect) BottomCenter() Vec2 {
	return Vec2{X: r.CenterX(), Y: float64(r.BottomY())}
}

func (r Rect) LeftCenter() Vec2 {
	return Vec2{X: float64(r.X), Y: r.CenterY()}
}

func (r Rect) Center() Vec2 {
	return Vec2{X: r.CenterX(), Y: r.CenterY()}
}

func (r Rect) GetRelativePoint(relativeX, relativeY float64) Vec2 {
	return Vec2{
		X: float64(r.X) + float64(r.W)*relativeX,
		Y: float64(r.Y) + float64(r.H)*relativeY,
	}
}

func (r Rect) Point(index int) Point {
	switch index {
	case 0:
		return r.TL()
	case 1:
		return r.TR()
	case 2:
		return r.BR()
	case 3:
		return r.BL()
	default:
		panic("geom.Rect.Point: index out of range")
	}
}

func (r Rect) Area() int {
	return r.W * r.H
}

func (r Rect) Perimeter() int {
	return (r.W + r.H) * 2
}

func (r Rect) HorizontalAspectRatio() float64 {
	if r.H == 0 {
		return 0
	}
	return float64(r.W) / float64(r.H)
}

func (r *Rect) Set(x, y, w, h int) *Rect {
	r.X = x
	r.Y = y
	r.W = w
	r.H = h
	return r
}

func (r Rect) With(x, y, w, h int) Rect {
	r.Set(x, y, w, h)
	return r
}

func (r *Rect) SetRect(other Rect) *Rect {
	*r = other
	return r
}

func (r Rect) WithRect(other Rect) Rect {
	r.SetRect(other)
	return r
}

func (r *Rect) SetPos(pos Point) *Rect {
	r.X = pos.X
	r.Y = pos.Y
	return r
}

func (r Rect) WithPos(pos Point) Rect {
	r.SetPos(pos)
	return r
}

func (r *Rect) SetCenter(center Point) *Rect {
	r.X = center.X - r.W/2
	r.Y = center.Y - r.H/2
	return r
}

func (r Rect) WithCenter(center Point) Rect {
	r.SetCenter(center)
	return r
}

func (r *Rect) SetSize(size Point) *Rect {
	r.W = size.X
	r.H = size.Y
	return r
}

func (r Rect) WithSize(size Point) Rect {
	r.SetSize(size)
	return r
}

func (r Rect) MovedBy(x, y int) Rect {
	r.X += x
	r.Y += y
	return r
}

func (r Rect) MovedByPoint(delta Point) Rect {
	return r.MovedBy(delta.X, delta.Y)
}

func (r *Rect) MoveBy(x, y int) *Rect {
	r.X += x
	r.Y += y
	return r
}

func (r *Rect) MoveByPoint(delta Point) *Rect {
	return r.MoveBy(delta.X, delta.Y)
}

func (r Rect) Stretched(xy int) Rect {
	return r.StretchedXY(xy, xy)
}

func (r Rect) StretchedXY(x, y int) Rect {
	return Rect{
		X: r.X - x,
		Y: r.Y - y,
		W: r.W + x*2,
		H: r.H + y*2,
	}
}

func (r Rect) StretchedPoint(xy Point) Rect {
	return r.StretchedXY(xy.X, xy.Y)
}

func (r Rect) StretchedLTRB(top, right, bottom, left int) Rect {
	return Rect{
		X: r.X - left,
		Y: r.Y - top,
		W: r.W + left + right,
		H: r.H + top + bottom,
	}
}

func (r Rect) Rotated90(n int) Rect {
	if n%2 == 0 {
		return r
	}
	original := r
	result := EmptyRect()
	result.SetPos(original.BL().Rotated90At(rectCenterPoint(original), 1)).SetSize(original.Size().YX())
	return result
}

func (r *Rect) Rotate90(n int) *Rect {
	*r = r.Rotated90(n)
	return r
}

func (r Rect) Rotated90At(pos Point, n int) Rect {
	original := r
	result := EmptyRect()
	switch n % 4 {
	case 1, -3:
		result.SetPos(original.BL().Rotated90At(pos, 1)).SetSize(original.Size().YX())
		return result
	case 2, -2:
		result.SetPos(original.BR().Rotated90At(pos, 2)).SetSize(original.Size())
		return result
	case 3, -1:
		result.SetPos(original.TR().Rotated90At(pos, 3)).SetSize(original.Size().YX())
		return result
	default:
		return r
	}
}

func (r *Rect) Rotate90At(pos Point, n int) *Rect {
	*r = r.Rotated90At(pos, n)
	return r
}

func (r Rect) Lerp(other Rect, f float64) RectF {
	return RectF{
		X: float64(r.X) + float64(other.X-r.X)*f,
		Y: float64(r.Y) + float64(other.Y-r.Y)*f,
		W: float64(r.W) + float64(other.W-r.W)*f,
		H: float64(r.H) + float64(other.H-r.H)*f,
	}
}

func (r Rect) GetOverlap(other Rect) Rect {
	ox := max(r.X, other.X)
	oy := max(r.Y, other.Y)
	ow := min(r.RightX(), other.RightX()) - ox
	if ow >= 0 {
		oh := min(r.BottomY(), other.BottomY()) - oy
		if oh >= 0 {
			return Rect{X: ox, Y: oy, W: ow, H: oh}
		}
	}
	return EmptyRect()
}

func (r Rect) IntersectsRect(other Rect) bool {
	return r.RightX() >= other.X &&
		other.RightX() >= r.X &&
		r.BottomY() >= other.Y &&
		other.BottomY() >= r.Y
}

func (r Rect) Intersects(other Rect) bool {
	return r.IntersectsRect(other)
}

func (r Rect) ContainsPoint(point Point) bool {
	return r.X <= point.X && point.X <= r.RightX() &&
		r.Y <= point.Y && point.Y <= r.BottomY()
}

func (r Rect) Contains(point Point) bool {
	return r.ContainsPoint(point)
}

func (r Rect) ContainsRect(other Rect) bool {
	return r.X <= other.X &&
		other.RightX() <= r.RightX() &&
		r.Y <= other.Y &&
		other.BottomY() <= r.BottomY()
}

func (r Rect) AsRectF() RectF {
	return RectF{
		X: float64(r.X),
		Y: float64(r.Y),
		W: float64(r.W),
		H: float64(r.H),
	}
}

func rectCenterPoint(r Rect) Point {
	return Point{X: r.X + r.W/2, Y: r.Y + r.H/2}
}
