package main

import (
	"fmt"
	"math"
)

// Функция для вычисления изменений dr/dt и df/dt
func derivatives(r, f, alpha float64) (float64, float64) {
	dr := 2*r - alpha*r*f
	df := -f + alpha*r*f
	return dr, df
}

// Метод Рунге-Кутты 4-го порядка
func rungeKutta(r0, f0, alpha, dt, tMax float64) ([]float64, []float64, []float64) {
	n := int(math.Ceil(tMax / dt))
	t := make([]float64, n)
	r := make([]float64, n)
	f := make([]float64, n)

	// Устанавливаем начальные условия
	r[0] = r0
	f[0] = f0
	t[0] = 0

	for i := 1; i < n; i++ {
		t[i] = t[i-1] + dt

		// Промежуточные вычисления по Рунге-Кутте
		dr1, df1 := derivatives(r[i-1], f[i-1], alpha)
		dr2, df2 := derivatives(r[i-1]+0.5*dt*dr1, f[i-1]+0.5*dt*df1, alpha)
		dr3, df3 := derivatives(r[i-1]+0.5*dt*dr2, f[i-1]+0.5*dt*df2, alpha)
		dr4, df4 := derivatives(r[i-1]+dt*dr3, f[i-1]+dt*df3, alpha)

		// Обновляем значения r и f
		r[i] = r[i-1] + dt*(dr1+2*dr2+2*dr3+dr4)/6
		f[i] = f[i-1] + dt*(df1+2*df2+2*df3+df4)/6
	}

	return t, r, f
}

func Task11_2A() {
	// Параметры задачи
	alpha := 0.01
	dt := 0.1
	tMax := 200.0

	// Начальные условия
	initialConditions := []struct {
		r0 float64
		f0 float64
	}{
		{50, 5},
		{100, 10},
		{200, 20},
		{1000, 50},
	}

	fmt.Println("Динамика Лисов и Кроликов")
	for _, cond := range initialConditions {
		fmt.Printf("\nНачальные значения: r0=%.0f, f0=%.0f\n", cond.r0, cond.f0)

		// Решение системы уравнений
		t, r, f := rungeKutta(cond.r0, cond.f0, alpha, dt, tMax)

		// Вывод результатов
		for i := 0; i < len(t); i += 50 { // Печатаем каждую 50-ю точку для наглядности
			fmt.Printf("t=%.1f, Кролики=%.2f, Лисы=%.2f\n", t[i], r[i], f[i])
		}
	}
}

// func Task11_2B() {
// 	// Параметры задачи
// 	alpha := 0.01
// 	dt := 0.1
// 	tMax := 10.0

// 	// Случай 1: Вымирание кроликов
// 	fmt.Println("\nСлучай 1: Вымирание кроликов (r0=15, f0=22)")
// 	t, r, f := rungeKutta(15, 22, alpha, dt, tMax)
// 	for i := 0; i < len(t); i += 1 {
// 		fmt.Printf("t=%.1f, Кролики=%.2f, Лисы=%.2f\n", t[i], r[i], f[i])
// 	}

// 	// Случай 2: Вымирание лис
// 	fmt.Println("\nСлучай 2: Вымирание лис (r0=100, f0=1)")
// 	t, r, f = rungeKutta(100, 1, alpha, dt, tMax)
// 	for i := 0; i < len(t); i += 1 {
// 		fmt.Printf("t=%.1f, Кролики=%.2f, Лисы=%.2f\n", t[i], r[i], f[i])
// 	}

// 	// Случай 3: Вымирание обоих видов
// 	fmt.Println("\nСлучай 3: Вымирание обоих видов (r0=f0=10)")
// 	t, r, f = rungeKutta(10, 10, alpha, dt, tMax)
// 	for i := 0; i < len(t); i += 1 {
// 		fmt.Printf("t=%.1f, Кролики=%.2f, Лисы=%.2f\n", t[i], r[i], f[i])
// 	}
// }
