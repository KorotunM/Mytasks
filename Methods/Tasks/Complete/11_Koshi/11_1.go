package main

import (
	"fmt"
	"math"
	"time"

	"tasks/tasks/pkg" // Путь для вызова Task10_erf
)

// Вычисление значений erf(x) через дифференциальное уравнение
func computeErfByODE(step float64, maxX float64) []float64 {
	n := int(maxX/step) + 1
	erfValues := make([]float64, n)
	x := 0.0
	y := 0.0

	for i := 0; i < n; i++ {
		erfValues[i] = y
		y += step * (2 / math.Sqrt(math.Pi) * math.Exp(-x*x))
		x += step
	}

	return erfValues
}

// Сравнение времени выполнения методов
func compareTime(step float64, maxX float64) {
	// Время для метода через ОДУ
	startODE := time.Now()
	erfODE := computeErfByODE(step, maxX)
	durationODE := time.Since(startODE)

	fmt.Println("Результаты через ОДУ:")
	for i, value := range erfODE {
		fmt.Printf("x=%.1f, erf(x)=%.6f\n", float64(i)*step, value)
	}
	fmt.Printf("Время выполнения метода ОДУ: %v\n", durationODE)

	// Время для метода из ЛР6
	startLR6 := time.Now()
	pkg.Task9_1() // Выводит данные из ЛР6
	durationLR6 := time.Since(startLR6)

	fmt.Printf("Время выполнения метода ЛР6: %v\n", durationLR6)
}

func Task11_1() {
	step := 0.1
	maxX := 2.0

	fmt.Println("Сравнение методов для вычисления erf(x):")
	compareTime(step, maxX)
}
