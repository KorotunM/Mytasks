package pkg

import (
	"fmt"
	"math"
)

// Функция для численного интегрирования с использованием метода трапеций
func trapezoidalRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

// Функция для вычисления erf(x) численно
func numericalErf(x float64) float64 {
	f := func(t float64) float64 {
		return math.Exp(-t * t)
	}
	integral := trapezoidalRule(f, 0, x, 1000) // Используем 1000 шагов для точности
	return (2 / math.Sqrt(math.Pi)) * integral
}

func Task9_1() {
	fmt.Printf("\nРезультаты через интегральное исчисление:\n")
	// Диапазон значений x, для которых будем вычислять erf(x)
	xValues := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0}

	// Вычисление значений функции erf(x) численно
	numericalValues := make([]float64, len(xValues))
	for i, x := range xValues {
		numericalValues[i] = numericalErf(x)
	}

	// Известные значения функции ошибок из math.Erf
	knownValues := make([]float64, len(xValues))
	for i, x := range xValues {
		knownValues[i] = math.Erf(x)
	}

	// Печать таблицы значений
	fmt.Printf("%-10s %-20s %-20s %-20s\n", "x", "Numerical erf(x)", "Known erf(x)", "Absolute Error")
	for i, x := range xValues {
		absoluteError := math.Abs(numericalValues[i] - knownValues[i])
		fmt.Printf("%-10.1f %-20.10f %-20.10f %-20.10f\n", x, numericalValues[i], knownValues[i], absoluteError)
	}

	// Суммарное отклонение по модулю
	totalError := 0.0
	for i := range xValues {
		totalError += math.Abs(numericalValues[i] - knownValues[i])
	}
	fmt.Printf("\nTotal Absolute Error: %f\n", totalError)
}
