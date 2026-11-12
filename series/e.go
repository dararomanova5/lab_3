package series

func E(eps float64) Result {
	r := Exp(1, eps)
	r.Name = "e"
	return r
}
