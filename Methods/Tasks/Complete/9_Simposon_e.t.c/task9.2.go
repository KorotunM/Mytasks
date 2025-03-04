package main

import (
	"fmt"
	"math"
)

// Функция для численного интегрирования с использованием метода трапеций
func trapezoidalRule_2(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

// Функция для численного интегрирования с использованием метода прямоугольников
func rectangleRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += f(a + float64(i)*h)
	}
	return sum * h
}

// Функция для численного интегрирования с использованием сплайн-квадратур (кубический сплайн)
func splineRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	xValues := make([]float64, n+1)
	yValues := make([]float64, n+1)

	for i := 0; i <= n; i++ {
		x := a + float64(i)*h
		xValues[i] = x
		yValues[i] = f(x)
	}

	// Вычисляем интеграл с использованием кусочно-кубического сплайна
	integral := 0.0
	for i := 0; i < n; i++ {
		h := xValues[i+1] - xValues[i]
		integral += (h / 6) * (yValues[i] + 4*f((xValues[i]+xValues[i+1])/2) + yValues[i+1])
	}

	return integral
}

// Функция для вычисления приближения числа Pi
func approximatePi(n int, method string) float64 {
	f := func(x float64) float64 {
		return 4 / (1 + x*x)
	}
	if method == "trapezoidal" {
		return trapezoidalRule_2(f, 0, 1, n)
	} else if method == "rectangle" {
		return rectangleRule(f, 0, 1, n)
	} else if method == "spline" {
		return splineRule(f, 0, 1, n)
	}
	return 0.0
}

func Task9_2() {
	fmt.Printf("\n Задание 2: \n")
	// Значения n для вычисления h = 1/n
	nValues := []int{8, 32, 128}

	fmt.Println("Using Trapezoidal Rule to approximate Pi:")
	for _, n := range nValues {
		h := 1.0 / float64(n)
		approxPi := approximatePi(n, "trapezoidal")
		error := math.Abs(math.Pi - approxPi)
		fmt.Printf("n = %d, h = %.5f, Approximation = %.10f, Error = %.10f\n", n, h, approxPi, error)
	}

	fmt.Println("\nUsing Rectangle Rule to approximate Pi:")
	for _, n := range nValues {
		h := 1.0 / float64(n)
		approxPi := approximatePi(n, "rectangle")
		error := math.Abs(math.Pi - approxPi)
		fmt.Printf("n = %d, h = %.5f, Approximation = %.10f, Error = %.10f\n", n, h, approxPi, error)
	}
	//для кубического сплайна ошибка пропорциональна примерно h^4
	fmt.Println("\nUsing Spline Rule to approximate Pi:")
	for _, n := range nValues {
		h := 1.0 / float64(n)
		approxPi := approximatePi(n, "spline")
		error := math.Abs(math.Pi - approxPi)
		fmt.Printf("n = %d, h = %.5f, Approximation = %.10f, Error = %.10f\n", n, h, approxPi, error)
	}
}
