package geom

// RectF is a rectangle with float64 components.
type RectF struct {
	X float64
	Y float64
	W float64
	H float64
}

func NewRectF(x, y, w, h float64) RectF {
	return RectF{X: x, Y: y, W: w, H: h}
}

func EmptyRectF() RectF {
	return RectF{}
}

func RectFromVec(a, b Vec2) RectF {
	x0, x1 := min(a.X, b.X), max(a.X, b.X)
	y0, y1 := min(a.Y, b.Y), max(a.Y, b.Y)
	return RectF{X: x0, Y: y0, W: x1 - x0, H: y1 - y0}
}

func (r RectF) Pos() Vec2 {
	return Vec2{X: r.X, Y: r.Y}
}

func (r RectF) Size() Vec2 {
	return Vec2{X: r.W, Y: r.H}
}

func (r RectF) IsEmpty() bool {
	return r.W == 0 || r.H == 0
}

func (r RectF) HasArea() bool {
	return r.W != 0 && r.H != 0
}

func (r RectF) LeftX() float64 {
	return r.X
}

func (r RectF) RightX() float64 {
	return r.X + r.W
}

func (r RectF) TopY() float64 {
	return r.Y
}

func (r RectF) BottomY() float64 {
	return r.Y + r.H
}

func (r RectF) CenterX() float64 {
	return r.X + r.W*0.5
}

func (r RectF) CenterY() float64 {
	return r.Y + r.H*0.5
}

func (r RectF) TL() Vec2 {
	return r.Pos()
}

func (r RectF) TR() Vec2 {
	return Vec2{X: r.RightX(), Y: r.Y}
}

func (r RectF) BR() Vec2 {
	return Vec2{X: r.RightX(), Y: r.BottomY()}
}

func (r RectF) BL() Vec2 {
	return Vec2{X: r.X, Y: r.BottomY()}
}

func (r RectF) TopCenter() Vec2 {
	return Vec2{X: r.CenterX(), Y: r.Y}
}

func (r RectF) RightCenter() Vec2 {
	return Vec2{X: r.RightX(), Y: r.CenterY()}
}

func (r RectF) BottomCenter() Vec2 {
	return Vec2{X: r.CenterX(), Y: r.BottomY()}
}

func (r RectF) LeftCenter() Vec2 {
	return Vec2{X: r.X, Y: r.CenterY()}
}

func (r RectF) Center() Vec2 {
	return Vec2{X: r.CenterX(), Y: r.CenterY()}
}

func (r RectF) GetRelativePoint(relativeX, relativeY float64) Vec2 {
	return Vec2{
		X: r.X + r.W*relativeX,
		Y: r.Y + r.H*relativeY,
	}
}

func (r RectF) Point(index int) Vec2 {
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
		panic("geom.RectF.Point: index out of range")
	}
}

func (r RectF) Area() float64 {
	return r.W * r.H
}

func (r RectF) Perimeter() float64 {
	return (r.W + r.H) * 2
}

func (r RectF) HorizontalAspectRatio() float64 {
	if r.H == 0 {
		return 0
	}
	return r.W / r.H
}

func (r *RectF) Set(x, y, w, h float64) *RectF {
	r.X = x
	r.Y = y
	r.W = w
	r.H = h
	return r
}

func (r RectF) With(x, y, w, h float64) RectF {
	r.Set(x, y, w, h)
	return r
}

func (r *RectF) SetRect(other RectF) *RectF {
	*r = other
	return r
}

func (r RectF) WithRect(other RectF) RectF {
	r.SetRect(other)
	return r
}

func (r *RectF) SetPos(pos Vec2) *RectF {
	r.X = pos.X
	r.Y = pos.Y
	return r
}

func (r RectF) WithPos(pos Vec2) RectF {
	r.SetPos(pos)
	return r
}

func (r *RectF) SetCenter(center Vec2) *RectF {
	r.X = center.X - r.W*0.5
	r.Y = center.Y - r.H*0.5
	return r
}

func (r RectF) WithCenter(center Vec2) RectF {
	r.SetCenter(center)
	return r
}

func (r *RectF) SetSize(size Vec2) *RectF {
	r.W = size.X
	r.H = size.Y
	return r
}

func (r RectF) WithSize(size Vec2) RectF {
	r.SetSize(size)
	return r
}

func (r RectF) MovedBy(x, y float64) RectF {
	r.X += x
	r.Y += y
	return r
}

func (r RectF) MovedByVec(delta Vec2) RectF {
	return r.MovedBy(delta.X, delta.Y)
}

func (r *RectF) MoveBy(x, y float64) *RectF {
	r.X += x
	r.Y += y
	return r
}

func (r *RectF) MoveByVec(delta Vec2) *RectF {
	return r.MoveBy(delta.X, delta.Y)
}

func (r RectF) Stretched(xy float64) RectF {
	return r.StretchedXY(xy, xy)
}

func (r RectF) StretchedXY(x, y float64) RectF {
	return RectF{
		X: r.X - x,
		Y: r.Y - y,
		W: r.W + x*2,
		H: r.H + y*2,
	}
}

func (r RectF) StretchedVec(xy Vec2) RectF {
	return r.StretchedXY(xy.X, xy.Y)
}

func (r RectF) StretchedLTRB(top, right, bottom, left float64) RectF {
	return RectF{
		X: r.X - left,
		Y: r.Y - top,
		W: r.W + left + right,
		H: r.H + top + bottom,
	}
}

func (r RectF) Scaled(s float64) RectF {
	return r.ScaledXY(s, s)
}

func (r RectF) ScaledXY(sx, sy float64) RectF {
	return EmptyRectF().WithSize(Vec2{X: r.W * sx, Y: r.H * sy}).WithCenter(r.Center())
}

func (r RectF) ScaledVec(s Vec2) RectF {
	return r.ScaledXY(s.X, s.Y)
}

func (r RectF) ScaledAt(pos Vec2, s float64) RectF {
	return r.ScaledAtXY(pos, s, s)
}

func (r RectF) ScaledAtXY(pos Vec2, sx, sy float64) RectF {
	return RectF{
		X: pos.X + (r.X-pos.X)*sx,
		Y: pos.Y + (r.Y-pos.Y)*sy,
		W: r.W * sx,
		H: r.H * sy,
	}
}

func (r RectF) ScaledAtVec(pos, s Vec2) RectF {
	return r.ScaledAtXY(pos, s.X, s.Y)
}

func (r RectF) Rotated90(n int) RectF {
	if n%2 == 0 {
		return r
	}
	return EmptyRectF().WithPos(r.BL().Rotated90At(r.Center(), 1)).WithSize(r.Size().YX())
}

func (r *RectF) Rotate90(n int) *RectF {
	*r = r.Rotated90(n)
	return r
}

func (r RectF) Rotated90At(pos Vec2, n int) RectF {
	switch n % 4 {
	case 1, -3:
		return EmptyRectF().WithPos(r.BL().Rotated90At(pos, 1)).WithSize(r.Size().YX())
	case 2, -2:
		return EmptyRectF().WithPos(r.BR().Rotated90At(pos, 2)).WithSize(r.Size())
	case 3, -1:
		return EmptyRectF().WithPos(r.TR().Rotated90At(pos, 3)).WithSize(r.Size().YX())
	default:
		return r
	}
}

func (r *RectF) Rotate90At(pos Vec2, n int) *RectF {
	*r = r.Rotated90At(pos, n)
	return r
}

func (r RectF) Lerp(other RectF, f float64) RectF {
	return RectF{
		X: r.X + (other.X-r.X)*f,
		Y: r.Y + (other.Y-r.Y)*f,
		W: r.W + (other.W-r.W)*f,
		H: r.H + (other.H-r.H)*f,
	}
}

func (r RectF) GetOverlap(other RectF) RectF {
	ox := max(r.X, other.X)
	oy := max(r.Y, other.Y)
	ow := min(r.RightX(), other.RightX()) - ox
	if ow >= 0 {
		oh := min(r.BottomY(), other.BottomY()) - oy
		if oh >= 0 {
			return RectF{X: ox, Y: oy, W: ow, H: oh}
		}
	}
	return EmptyRectF()
}

func (r RectF) IntersectsRect(other RectF) bool {
	return r.RightX() >= other.X &&
		other.RightX() >= r.X &&
		r.BottomY() >= other.Y &&
		other.BottomY() >= r.Y
}

func (r RectF) Intersects(other RectF) bool {
	return r.IntersectsRect(other)
}

func (r RectF) ContainsPoint(point Vec2) bool {
	return r.X <= point.X && point.X <= r.RightX() &&
		r.Y <= point.Y && point.Y <= r.BottomY()
}

func (r RectF) Contains(point Vec2) bool {
	return r.ContainsPoint(point)
}

func (r RectF) ContainsRect(other RectF) bool {
	return r.X <= other.X &&
		other.RightX() <= r.RightX() &&
		r.Y <= other.Y &&
		other.BottomY() <= r.BottomY()
}

func (r RectF) AsRect() Rect {
	return Rect{
		X: int(r.X),
		Y: int(r.Y),
		W: int(r.W),
		H: int(r.H),
	}
}
