package main

import (
	"fmt"
	"math"
)

// Функция для численного интегрирования с использованием метода Симпсона
func simpsonsRule(f func(float64) float64, a, b float64, n int) float64 {
	// n должен быть четным для метода Симпсона
	if n%2 != 0 {
		n++
	}

	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}

	return (h / 3) * sum
}

func Task9_3() {
	// Определение функции f(x)
	fmt.Printf("\n Задание 3: \n")
	f := func(x float64) float64 {
		if x >= 0 && x <= 2 {
			return math.Exp(x * x)
		} else if x > 2 && x <= 4 {
			return 1 / (4 - math.Sin(16*math.Pi*x))
		}
		return 0
	}

	// Вычисление интеграла на интервале [0, 4] методом Симпсона
	n := 1000 // Количество разбиений, должно быть четным
	integral := simpsonsRule(f, 0, 4, n)

	// Вывод результата
	fmt.Printf("Approximate value of the integral: %.10f\n", integral)
}
