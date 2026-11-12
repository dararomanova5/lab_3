package series

import (
	"fmt"
	"math"
)

type StopReason int

const (
	StopConverged StopReason = iota
	StopMaxIter
	StopNaN
)

func (s StopReason) String() string {
	switch s {
	case StopConverged:
		return "точность"
	case StopMaxIter:
		return "лимит итераций"
	case StopNaN:
		return "NaN/Inf"
	}
	return "?"
}

type Result struct {
	Name       string
	Approx     float64
	Exact      float64
	Iterations int
	Stop       StopReason
}

func (r Result) Converged() bool { return r.Stop == StopConverged }
func (r Result) Error() float64  { return math.Abs(r.Approx - r.Exact) }

func (r Result) String() string {
	return fmt.Sprintf("%-20s приближ=%.12f точное=%.12f погр=%.2e итер=%d [%s]",
		r.Name, r.Approx, r.Exact, r.Error(), r.Iterations, r.Stop)
}

func badReport(name string, approx, exact float64, it int) Result {
	return Result{name, approx, exact, it, StopNaN}
}
