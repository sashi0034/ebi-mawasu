package geom

import "math"

// Vec2 is a 2D vector with float64 components.
type Vec2 struct {
	X float64
	Y float64
}

func V2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func Zero() Vec2 {
	return Vec2{}
}

func One() Vec2 {
	return Vec2{X: 1, Y: 1}
}

func All(value float64) Vec2 {
	return Vec2{X: value, Y: value}
}

func UnitX() Vec2 {
	return Vec2{X: 1}
}

func UnitY() Vec2 {
	return Vec2{Y: 1}
}

func Left(length float64) Vec2 {
	return Vec2{X: -length}
}

func Right(length float64) Vec2 {
	return Vec2{X: length}
}

func Up(length float64) Vec2 {
	return Vec2{Y: -length}
}

func Down(length float64) Vec2 {
	return Vec2{Y: length}
}

func AnchorCenter() Vec2 {
	return Vec2{X: 0.5, Y: 0.5}
}

func AnchorTopLeft() Vec2 {
	return Vec2{}
}

func AnchorTopCenter() Vec2 {
	return Vec2{X: 0.5}
}

func AnchorTopRight() Vec2 {
	return Vec2{X: 1}
}

func AnchorRightCenter() Vec2 {
	return Vec2{X: 1, Y: 0.5}
}

func AnchorBottomRight() Vec2 {
	return Vec2{X: 1, Y: 1}
}

func AnchorBottomCenter() Vec2 {
	return Vec2{X: 0.5, Y: 1}
}

func AnchorBottomLeft() Vec2 {
	return Vec2{Y: 1}
}

func AnchorLeftCenter() Vec2 {
	return Vec2{Y: 0.5}
}

func (v Vec2) Elem(index int) float64 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	default:
		return 0
	}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{X: v.X - other.X, Y: v.Y - other.Y}
}

func (v Vec2) Mul(s float64) Vec2 {
	return Vec2{X: v.X * s, Y: v.Y * s}
}

func (v Vec2) MulElem(other Vec2) Vec2 {
	return Vec2{X: v.X * other.X, Y: v.Y * other.Y}
}

func (v Vec2) Div(s float64) Vec2 {
	return v.Mul(1 / s)
}

func (v Vec2) DivElem(other Vec2) Vec2 {
	return Vec2{X: v.X / other.X, Y: v.Y / other.Y}
}

func (v Vec2) Neg() Vec2 {
	return Vec2{X: -v.X, Y: -v.Y}
}

func (v Vec2) EpsilonEquals(other Vec2, epsilon float64) bool {
	return math.Abs(v.X-other.X) <= epsilon &&
		math.Abs(v.Y-other.Y) <= epsilon
}

func (v Vec2) HasSameDirection(other Vec2) bool {
	return v.Dot(other) > 0
}

func (v Vec2) HasOppositeDirection(other Vec2) bool {
	return v.Dot(other) < 0
}

func (v Vec2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec2) HasNaN() bool {
	return math.IsNaN(v.X) || math.IsNaN(v.Y)
}

func (v Vec2) MinComponent() float64 {
	return min(v.X, v.Y)
}

func (v Vec2) MaxComponent() float64 {
	return max(v.X, v.Y)
}

func (v *Vec2) Clear() {
	v.X = 0
	v.Y = 0
}

func (v Vec2) WithX(x float64) Vec2 {
	return Vec2{X: x, Y: v.Y}
}

func (v Vec2) WithY(y float64) Vec2 {
	return Vec2{X: v.X, Y: y}
}

func (v *Vec2) Set(x, y float64) *Vec2 {
	v.X = x
	v.Y = y
	return v
}

func (v Vec2) With(x, y float64) Vec2 {
	v.Set(x, y)
	return v
}

func (v *Vec2) SetVec(other Vec2) *Vec2 {
	*v = other
	return v
}

func (v Vec2) WithVec(other Vec2) Vec2 {
	v.SetVec(other)
	return v
}

func (v Vec2) MovedBy(x, y float64) Vec2 {
	return Vec2{X: v.X + x, Y: v.Y + y}
}

func (v Vec2) MovedByVec(delta Vec2) Vec2 {
	return v.Add(delta)
}

func (v *Vec2) MoveBy(x, y float64) *Vec2 {
	v.X += x
	v.Y += y
	return v
}

func (v *Vec2) MoveByVec(delta Vec2) *Vec2 {
	return v.MoveBy(delta.X, delta.Y)
}

func (v Vec2) Clamped(rect RectF) Vec2 {
	return Vec2{
		X: min(max(v.X, rect.LeftX()), rect.RightX()),
		Y: min(max(v.Y, rect.TopY()), rect.BottomY()),
	}
}

func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Cross(other Vec2) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v Vec2) HorizontalAspectRatio() float64 {
	if v.Y == 0 {
		return 0
	}
	return v.X / v.Y
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v Vec2) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) InvLength() float64 {
	return 1 / v.Length()
}

func (v Vec2) ManhattanLength() float64 {
	return math.Abs(v.X) + math.Abs(v.Y)
}

func (v Vec2) ManhattanDistanceFrom(other Vec2) float64 {
	return math.Abs(v.X-other.X) + math.Abs(v.Y-other.Y)
}

func (v Vec2) DistanceFrom(other Vec2) float64 {
	return math.Sqrt(v.DistanceFromSq(other))
}

func (v Vec2) DistanceFromSq(other Vec2) float64 {
	x := v.X - other.X
	y := v.Y - other.Y
	return x*x + y*y
}

func (v Vec2) WithLength(length float64) Vec2 {
	if current := v.Length(); current != 0 {
		return v.Mul(length / current)
	}
	return v
}

func (v *Vec2) SetLength(length float64) *Vec2 {
	*v = v.WithLength(length)
	return v
}

func (v Vec2) LimitLength(maxLength float64) Vec2 {
	if length := v.Length(); length > maxLength {
		return v.Mul(maxLength / length)
	}
	return v
}

func (v *Vec2) LimitLengthSelf(maxLength float64) *Vec2 {
	*v = v.LimitLength(maxLength)
	return v
}

func (v Vec2) Normalized() Vec2 {
	lengthSq := v.LengthSq()
	if lengthSq == 0 {
		return v
	}
	return v.Mul(1 / math.Sqrt(lengthSq))
}

func (v *Vec2) Normalize() *Vec2 {
	*v = v.Normalized()
	return v
}

func (v Vec2) NormalizedOr(valueIfZero Vec2) Vec2 {
	lengthSq := v.LengthSq()
	if lengthSq == 0 {
		return valueIfZero
	}
	return v.Mul(1 / math.Sqrt(lengthSq))
}

func (v Vec2) Rotated(angle float64) Vec2 {
	s, c := math.Sincos(angle)
	return Vec2{
		X: v.X*c - v.Y*s,
		Y: v.X*s + v.Y*c,
	}
}

func (v Vec2) RotatedAt(center Vec2, angle float64) Vec2 {
	return v.Sub(center).Rotated(angle).Add(center)
}

func (v *Vec2) Rotate(angle float64) *Vec2 {
	*v = v.Rotated(angle)
	return v
}

func (v *Vec2) RotateAt(center Vec2, angle float64) *Vec2 {
	*v = v.RotatedAt(center, angle)
	return v
}

func (v Vec2) Rotated90(n int) Vec2 {
	switch n % 4 {
	case 1, -3:
		return Vec2{X: -v.Y, Y: v.X}
	case 2, -2:
		return Vec2{X: -v.X, Y: -v.Y}
	case 3, -1:
		return Vec2{X: v.Y, Y: -v.X}
	default:
		return v
	}
}

func (v Vec2) Rotated90At(center Vec2, n int) Vec2 {
	return v.Sub(center).Rotated90(n).Add(center)
}

func (v *Vec2) Rotate90(n int) *Vec2 {
	*v = v.Rotated90(n)
	return v
}

func (v *Vec2) Rotate90At(center Vec2, n int) *Vec2 {
	*v = v.Rotated90At(center, n)
	return v
}

func (v Vec2) Angle() float64 {
	if v.IsZero() {
		return math.NaN()
	}
	return math.Atan2(v.X, -v.Y)
}

func (v Vec2) GetAngle() float64 {
	return v.Angle()
}

func (v Vec2) AngleTo(other Vec2) float64 {
	if v.IsZero() || other.IsZero() {
		return math.NaN()
	}
	return math.Atan2(v.Cross(other), v.Dot(other))
}

func (v Vec2) GetAngleTo(other Vec2) float64 {
	return v.AngleTo(other)
}

func (v Vec2) PerpendicularCW() Vec2 {
	return Vec2{X: -v.Y, Y: v.X}
}

func (v Vec2) GetPerpendicularCW() Vec2 {
	return v.PerpendicularCW()
}

func (v Vec2) PerpendicularCCW() Vec2 {
	return Vec2{X: v.Y, Y: -v.X}
}

func (v Vec2) GetPerpendicularCCW() Vec2 {
	return v.PerpendicularCCW()
}

func (v Vec2) Midpoint(other Vec2) Vec2 {
	return Vec2{X: (v.X + other.X) * 0.5, Y: (v.Y + other.Y) * 0.5}
}

func (v Vec2) GetMidpoint(other Vec2) Vec2 {
	return v.Midpoint(other)
}

func (v Vec2) Projection(onto Vec2) Vec2 {
	if lengthSq := onto.LengthSq(); lengthSq != 0 {
		return onto.Mul(v.Dot(onto) / lengthSq)
	}
	return Vec2{}
}

func (v Vec2) PointByAngleAndDistance(angle, distance float64) Vec2 {
	s, c := math.Sincos(angle)
	return Vec2{X: s*distance + v.X, Y: -c*distance + v.Y}
}

func (v Vec2) GetPointByAngleAndDistance(angle, distance float64) Vec2 {
	return v.PointByAngleAndDistance(angle, distance)
}

func (v Vec2) Lerp(other Vec2, f float64) Vec2 {
	return Vec2{
		X: v.X + (other.X-v.X)*f,
		Y: v.Y + (other.Y-v.Y)*f,
	}
}

func (v Vec2) XX() Vec2 {
	return Vec2{X: v.X, Y: v.X}
}

func (v Vec2) XY() Vec2 {
	return v
}

func (v Vec2) YX() Vec2 {
	return Vec2{X: v.Y, Y: v.X}
}

func (v Vec2) YY() Vec2 {
	return Vec2{X: v.Y, Y: v.Y}
}

func (v Vec2) X0() Vec2 {
	return Vec2{X: v.X}
}

func (v Vec2) Y0() Vec2 {
	return Vec2{X: v.Y}
}

func (v Vec2) AsPoint() Point {
	return Point{X: int(v.X), Y: int(v.Y)}
}
