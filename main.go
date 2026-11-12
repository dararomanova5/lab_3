package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"go-lab3/series"
)

// readFloat читает число, принимая '.' и ',' как разделитель,
// пробелы внутри ("1 234,5") и отбрасывая NaN/Inf.
func readFloat(r *bufio.Reader, prompt string, def float64) float64 {
	for {
		fmt.Print(prompt)
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения, использую по умолчанию:", def)
			return def
		}
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, ",", ".")
		line = strings.ReplaceAll(line, " ", "")
		if line == "" {
			fmt.Println("Пустой ввод, использую по умолчанию:", def)
			return def
		}
		v, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Не удалось разобрать %q, попробуйте ещё раз (пример: 1,5)\n", line)
			continue
		}
		if math.IsNaN(v) || math.IsInf(v, 0) {
			fmt.Println("NaN/Inf недопустимы, повторите ввод")
			continue
		}
		return v
	}
}

// readEps читает точность ε > 0.
func readEps(r *bufio.Reader) float64 {
	for {
		v := readFloat(r, "Введите точность ε (например 1e-8 или 0,00000001): ", 1e-8)
		if v <= 0 {
			fmt.Println("ε должно быть > 0, повторите")
			continue
		}
		return v
	}
}

// printTable выводит результаты в консоль.
func printTable(title string, rs []series.Result) {
	line := strings.Repeat("-", 96)
	fmt.Println(title)
	fmt.Println(line)
	fmt.Printf("%-20s | %-16s | %-16s | %-11s | %-7s | %s\n",
		"Вычисление", "Приближённое", "Точное", "Погрешность", "Итераций", "Статус")
	fmt.Println(line)
	for _, r := range rs {
		fmt.Printf("%-20s | %-16.12f | %-16.12f | %-11.2e | %-7d | %s\n",
			r.Name, r.Approx, r.Exact, r.Error(), r.Iterations, r.Stop)
	}
	fmt.Println(line)
	fmt.Println()
}

// ---------- Интеграция (общая) ----------

func main() {
	r := bufio.NewReader(os.Stdin)

	x := readFloat(r, "Введите x (например 1,5): ", 1.0)
	eps := readEps(r)
	fmt.Println()

	results := []series.Result{
		series.Sin(x, eps),
		series.Cos(x, eps),
		series.Exp(x, eps),
		series.E(eps),
		series.PiLeibniz(eps),
		series.PiNilakantha(eps),
		series.PiMachin(eps),
		series.SumOddSkip(eps),
	}

	printTable(fmt.Sprintf("Результаты при x=%g, ε=%g", x, eps), results)
	series.IterationsStudy()
}
