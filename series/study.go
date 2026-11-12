package series

import (
	"fmt"
	"strings"
)

func IterationsStudy() {
	epsList := []float64{1e-2, 1e-4, 1e-6, 1e-8, 1e-10}
	fmt.Println("Анализ числа итераций при разной точности ε")
	fmt.Println(strings.Repeat("-", 86))
	fmt.Printf("%-10s | %-10s | %-10s | %-14s | %-13s | %-10s\n",
		"ε", "sin(1)", "exp(1)", "π Нилаканта", "π Лейбниц", "π Мачин")
	fmt.Println(strings.Repeat("-", 86))
	for _, eps := range epsList {
		fmt.Printf("%-10.0e | %-10d | %-10d | %-14d | %-13d | %-10d\n",
			eps,
			Sin(1, eps).Iterations,
			Exp(1, eps).Iterations,
			PiNilakantha(eps).Iterations,
			PiLeibniz(eps).Iterations,
			PiMachin(eps).Iterations,
		)
	}
	fmt.Println()
}
