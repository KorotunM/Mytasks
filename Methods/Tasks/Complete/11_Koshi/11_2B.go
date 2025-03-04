package main

import (
	"fmt"
	"math"
)

// Функция для вычисления изменений dr/dt и df/dt
func derivativesStop(r, f, alpha float64) (float64, float64) {
	dr := 2*r - alpha*r*f
	df := -f + alpha*r*f
	return dr, df
}

// Метод Рунге-Кутты 4-го порядка
func rungeKuttaWithStop(r0, f0, alpha, dt, tMax float64, caseName string) {
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
		dr1, df1 := derivativesStop(r[i-1], f[i-1], alpha)
		dr2, df2 := derivativesStop(r[i-1]+0.5*dt*dr1, f[i-1]+0.5*dt*df1, alpha)
		dr3, df3 := derivativesStop(r[i-1]+0.5*dt*dr2, f[i-1]+0.5*dt*df2, alpha)
		dr4, df4 := derivativesStop(r[i-1]+dt*dr3, f[i-1]+dt*df3, alpha)

		// Обновляем значения r и f
		r[i] = r[i-1] + dt*(dr1+2*dr2+2*dr3+dr4)/6
		f[i] = f[i-1] + dt*(df1+2*df2+2*df3+df4)/6

		// Проверка условий вымирания
		if r[i] < 1 && f[i] >= 1 && caseName == "Кролики" {
			fmt.Printf("t=%.1f, Кролики вымерли! (Кролики=%.2f, Лисы=%.2f)\n", t[i], r[i], f[i])
			return
		}
		if f[i] < 1 && r[i] >= 1 && caseName == "Лисы" {
			fmt.Printf("t=%.1f, Лисы вымерли! (Кролики=%.2f, Лисы=%.2f)\n", t[i], r[i], f[i])
			return
		}
		if r[i] < 1 && f[i] < 1 && caseName == "Оба вида" {
			fmt.Printf("t=%.1f, Оба вида вымерли! (Кролики=%.2f, Лисы=%.2f)\n", t[i], r[i], f[i])
			return
		}
	}

	// Если вымирания не произошло
	fmt.Printf("t=%.1f, Итерации завершены без вымирания. (Кролики=%.2f, Лисы=%.2f)\n", tMax, r[n-1], f[n-1])
}

func Task11_2B() {
	// Параметры задачи
	alpha := 0.01
	dt := 0.01
	tMax := 1000.0

	// Случай 1: Вымирание кроликов
	fmt.Println("\nСлучай 1: Вымирание кроликов (r0=15, f0=22)")
	rungeKuttaWithStop(15, 22, alpha, dt, tMax, "Кролики")

	// Случай 2: Вымирание лис
	fmt.Println("\nСлучай 2: Вымирание лис (r0=1000, f0=1)")
	rungeKuttaWithStop(1000, 1, alpha, dt, tMax, "Лисы")

	// Случай 3: Вымирание обоих видов
	fmt.Println("\nСлучай 3: Вымирание обоих видов (r0=f0=5)")
	rungeKuttaWithStop(1000, 1000, alpha, dt, tMax, "Оба вида")
}
