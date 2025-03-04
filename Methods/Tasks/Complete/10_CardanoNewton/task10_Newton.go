package main

import (
	"fmt"
	"math"
)

// Кубическое уравнение
func cubicEquation(x, alpha float64) float64 {
	return math.Pow(x, 3) + 3*math.Pow(x, 2) + alpha*alpha*x + 3*alpha*alpha
}

// Производная кубического уравнения
func cubicDerivative(x, alpha float64) float64 {
	return 3*math.Pow(x, 2) + 6*x + alpha*alpha
}

// Вторая производная кубического уравнения
func cubicSecondDerivative(x float64) float64 {
	return 6*x + 6
}

// Оценка погрешности метода Ньютона
func newtonError(x0, alpha float64, a, b float64) float64 {
	// Найти минимальное значение |F'(x)| на отрезке [a, b]
	m := math.Inf(1)
	for x := a; x <= b; x += 0.01 { // Шаг итерации 0.01
		absDerivative := math.Abs(cubicDerivative(x, alpha))
		if absDerivative < m {
			m = absDerivative
		}
	}

	// Вычислить ошибку по формуле
	fx0 := math.Abs(cubicEquation(x0, alpha))
	return fx0 / m
}

// Метод Ньютона
func newtonMethod(alpha float64, a, b float64, tolerance float64, maxIterations int) (float64, int, error) {
	// Выбор начального приближения
	var x0 float64
	if cubicEquation(a, alpha)*cubicSecondDerivative(a) > 0 {
		x0 = a
	} else if cubicEquation(b, alpha)*cubicSecondDerivative(b) > 0 {
		x0 = b
	} else {
		return 0, 0, fmt.Errorf("no suitable initial guess found in [%f, %f]", a, b)
	}

	// Итерации метода Ньютона
	x := x0
	for i := 0; i < maxIterations; i++ {
		fx := cubicEquation(x, alpha)
		fpx := cubicDerivative(x, alpha)

		if math.Abs(fpx) < 1e-15 {
			return x, i, fmt.Errorf("derivative is too small at iteration %d", i)
		}

		xNext := x - fx/fpx

		if math.Abs(xNext-x) < tolerance {
			return xNext, i + 1, nil
		}

		x = xNext
	}

	return x, maxIterations, fmt.Errorf("method did not converge within the maximum number of iterations")
}

func Task10_Newton() {
	// Значения alpha
	alphaValues := []float64{0.1, 5, 10, 113, 1000, 1e6, 1e9, 1e16, 1e18}

	// Интервал для начального приближения
	a := -10.0
	b := 10.0
	tolerance := 1e-9
	maxIterations := 100

	fmt.Println("=== Исследование метода Ньютона ===")
	fmt.Println("Alpha\t\tRoot (Num)\t\tIterations\tError Estimate")

	for _, alpha := range alphaValues {
		root, iterations, err := newtonMethod(alpha, a, b, tolerance, maxIterations)

		if err != nil {
			fmt.Printf("Alpha: %.2e\tError: %v\n", alpha, err)
		} else {
			errorEstimate := newtonError(root, alpha, a, b)
			fmt.Printf("Alpha: %.2e\tRoot: %.13f\tIter: %d    \tError: %.2e\n",
				alpha, root, iterations, errorEstimate)
		}
	}
}
