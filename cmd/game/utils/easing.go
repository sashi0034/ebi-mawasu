package utils

import "math"

type EasingFunc func(t float64) float64

func Linear(t float64) float64 {
	return t
}

func Sine(t float64) float64 {
	return 1 - math.Cos(t*math.Pi/2)
}

func Quad(t float64) float64 {
	return t * t
}

func Cubic(t float64) float64 {
	return t * t * t
}

func Quart(t float64) float64 {
	return t * t * t * t
}

func Quint(t float64) float64 {
	return t * t * t * t * t
}

func Expo(t float64) float64 {
	if t == 0 {
		return 0
	}
	return math.Exp2(10 * (t - 1))
}

func Circ(t float64) float64 {
	return 1 - math.Sqrt(1-t*t)
}

func Back(t float64) float64 {
	return t * t * (2.70158*t - 1.70158)
}

func Elastic(t float64) float64 {
	return elasticAP(t, 1, 0.3)
}

func Bounce(t float64) float64 {
	t = 1 - t
	if t < 1/2.75 {
		return 1 - 7.5625*t*t
	}
	if t < 2/2.75 {
		t -= 1.5 / 2.75
		return 1 - (7.5625*t*t + 0.75)
	}
	if t < 2.5/2.75 {
		t -= 2.25 / 2.75
		return 1 - (7.5625*t*t + 0.9375)
	}
	t -= 2.625 / 2.75
	return 1 - (7.5625*t*t + 0.984375)
}

func EaseIn(f EasingFunc, t float64) float64 {
	return f(t)
}

func EaseInValue(f EasingFunc, start, end, t float64) float64 {
	return Lerp(start, end, EaseIn(f, t))
}

func EaseOut(f EasingFunc, t float64) float64 {
	return 1 - f(1-t)
}

func EaseOutValue(f EasingFunc, start, end, t float64) float64 {
	return Lerp(start, end, EaseOut(f, t))
}

func EaseInOut(f EasingFunc, t float64) float64 {
	if t < 0.5 {
		return f(2*t) * 0.5
	}
	return 0.5 + EaseOut(f, 2*t-1)*0.5
}

func EaseInOutValue(f EasingFunc, start, end, t float64) float64 {
	return Lerp(start, end, EaseInOut(f, t))
}

func Lerp(start, end, t float64) float64 {
	return start + (end-start)*t
}

func EaseInLinear(t float64) float64 {
	return Linear(t)
}

func EaseOutLinear(t float64) float64 {
	return EaseOut(Linear, t)
}

func EaseInOutLinear(t float64) float64 {
	return EaseInOut(Linear, t)
}

func EaseInSine(t float64) float64 {
	return Sine(t)
}

func EaseOutSine(t float64) float64 {
	return EaseOut(Sine, t)
}

func EaseInOutSine(t float64) float64 {
	return EaseInOut(Sine, t)
}

func EaseInQuad(t float64) float64 {
	return Quad(t)
}

func EaseOutQuad(t float64) float64 {
	return EaseOut(Quad, t)
}

func EaseInOutQuad(t float64) float64 {
	return EaseInOut(Quad, t)
}

func EaseInCubic(t float64) float64 {
	return Cubic(t)
}

func EaseOutCubic(t float64) float64 {
	return EaseOut(Cubic, t)
}

func EaseInOutCubic(t float64) float64 {
	return EaseInOut(Cubic, t)
}

func EaseInQuart(t float64) float64 {
	return Quart(t)
}

func EaseOutQuart(t float64) float64 {
	return EaseOut(Quart, t)
}

func EaseInOutQuart(t float64) float64 {
	return EaseInOut(Quart, t)
}

func EaseInQuint(t float64) float64 {
	return Quint(t)
}

func EaseOutQuint(t float64) float64 {
	return EaseOut(Quint, t)
}

func EaseInOutQuint(t float64) float64 {
	return EaseInOut(Quint, t)
}

func EaseInExpo(t float64) float64 {
	return Expo(t)
}

func EaseOutExpo(t float64) float64 {
	return EaseOut(Expo, t)
}

func EaseInOutExpo(t float64) float64 {
	return EaseInOut(Expo, t)
}

func EaseInCirc(t float64) float64 {
	return Circ(t)
}

func EaseOutCirc(t float64) float64 {
	return EaseOut(Circ, t)
}

func EaseInOutCirc(t float64) float64 {
	return EaseInOut(Circ, t)
}

func EaseInBack(t float64) float64 {
	return Back(t)
}

func EaseOutBack(t float64) float64 {
	return EaseOut(Back, t)
}

func EaseInOutBack(t float64) float64 {
	return EaseInOut(Back, t)
}

func EaseInElastic(t float64) float64 {
	return Elastic(t)
}

func EaseOutElastic(t float64) float64 {
	return EaseOut(Elastic, t)
}

func EaseInOutElastic(t float64) float64 {
	return EaseInOut(Elastic, t)
}

func EaseInBounce(t float64) float64 {
	return Bounce(t)
}

func EaseOutBounce(t float64) float64 {
	return EaseOut(Bounce, t)
}

func EaseInOutBounce(t float64) float64 {
	return EaseInOut(Bounce, t)
}

func elasticAP(t, a, p float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return 1
	}

	var s float64
	if a < 1 {
		a = 1
		s = p / 4
	} else {
		s = p / (2 * math.Pi) * math.Asin(1/a)
	}

	t -= 1
	return -(a * math.Exp2(10*t) * math.Sin((t-s)*2*math.Pi/p))
}
