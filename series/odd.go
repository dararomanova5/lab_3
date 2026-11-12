package series

import "math"

func SumOddSkip(eps float64) Result {
	sum := 0.0
	iters := 0
	for k := 1; k <= maxIter; k++ {
		if k%2 == 0 {
			continue
		}
		iters++
		sum += 1 / float64(k*k)
		if math.Abs(sum-math.Pi*math.Pi/8) < eps {
			return Result{"Σ1/k² (нечёт.)", sum, math.Pi * math.Pi / 8, iters, StopConverged}
		}
	}
	return Result{"Σ1/k² (нечёт.)", sum, math.Pi * math.Pi / 8, iters, StopMaxIter}
}
