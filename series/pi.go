package series

import "math"

const maxIterLeibniz = 5_000_000

func PiLeibniz(eps float64) Result {
	sum := 0.0
	sign := 1.0
	for i := 0; i < maxIterLeibniz; i++ {
		sum += sign / float64(2*i+1)
		sign = -sign
		if math.Abs(4*sum-math.Pi) < eps {
			return Result{"π (Лейбниц)", 4 * sum, math.Pi, i + 1, StopConverged}
		}
	}
	return Result{"π (Лейбниц)", 4 * sum, math.Pi, maxIterLeibniz, StopMaxIter}
}

func PiNilakantha(eps float64) Result {
	pi := 3.0
	sign := 1.0
	for i := 1; i <= maxIter; i++ {
		a := float64(2 * i)
		pi += sign * 4 / (a * (a + 1) * (a + 2))
		sign = -sign
		if math.IsNaN(pi) || math.IsInf(pi, 0) {
			return badReport("π (Нилаканта)", pi, math.Pi, i)
		}
		if math.Abs(pi-math.Pi) < eps {
			return Result{"π (Нилаканта)", pi, math.Pi, i, StopConverged}
		}
	}
	return Result{"π (Нилаканта)", pi, math.Pi, maxIter, StopMaxIter}
}

func atanInv(q int, eps float64) (float64, int, bool) {
	qf := float64(q)
	q2 := qf * qf
	term := 1.0 / qf
	sum := 0.0
	sign := 1.0
	for n := 0; n <= maxIter; n++ {
		part := sign * term / float64(2*n+1)
		if math.Abs(part) < eps {
			return sum, n, true
		}
		sum += part
		term /= q2
		sign = -sign
	}
	return sum, maxIter, false
}

func PiMachin(eps float64) Result {
	a, ia, oka := atanInv(5, eps)
	b, ib, okb := atanInv(239, eps)
	pi := 4 * (4*a - b)
	iters := ia + ib
	if oka && okb {
		return Result{"π (Мачин)", pi, math.Pi, iters, StopConverged}
	}
	return Result{"π (Мачин)", pi, math.Pi, iters, StopMaxIter}
}
