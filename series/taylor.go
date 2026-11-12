package series

import "math"

const maxIter = 10_000_000

func reduceAngle(x float64) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}
	x = math.Mod(x, 2*math.Pi)
	if x > math.Pi {
		x -= 2 * math.Pi
	} else if x < -math.Pi {
		x += 2 * math.Pi
	}
	return x
}

func Sin(x, eps float64) Result {
	x = reduceAngle(x)
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return badReport("sin(x)", x, math.NaN(), 0)
	}
	term, sum := x, x
	for i := 1; i <= maxIter; i++ {
		term *= -x * x / float64((2*i)*(2*i+1))
		if math.IsNaN(term) || math.IsInf(term, 0) {
			return badReport("sin(x)", sum, math.Sin(x), i)
		}
		if math.Abs(term) < eps {
			return Result{"sin(x)", sum, math.Sin(x), i, StopConverged}
		}
		sum += term
	}
	return Result{"sin(x)", sum, math.Sin(x), maxIter, StopMaxIter}
}

func Cos(x, eps float64) Result {
	x = reduceAngle(x)
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return badReport("cos(x)", x, math.NaN(), 0)
	}
	term, sum := 1.0, 1.0
	for i := 1; i <= maxIter; i++ {
		term *= -x * x / float64((2*i-1)*(2*i))
		if math.IsNaN(term) || math.IsInf(term, 0) {
			return badReport("cos(x)", sum, math.Cos(x), i)
		}
		if math.Abs(term) < eps {
			return Result{"cos(x)", sum, math.Cos(x), i, StopConverged}
		}
		sum += term
	}
	return Result{"cos(x)", sum, math.Cos(x), maxIter, StopMaxIter}
}

func Exp(x, eps float64) Result {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return badReport("exp(x)", x, math.Exp(x), 0)
	}
	term, sum := 1.0, 1.0
	for i := 1; i <= maxIter; i++ {
		term *= x / float64(i)
		if math.IsNaN(term) || math.IsInf(term, 0) {
			return badReport("exp(x)", sum, math.Exp(x), i)
		}
		if math.Abs(term) < eps {
			return Result{"exp(x)", sum, math.Exp(x), i, StopConverged}
		}
		sum += term
	}
	return Result{"exp(x)", sum, math.Exp(x), maxIter, StopMaxIter}
}
