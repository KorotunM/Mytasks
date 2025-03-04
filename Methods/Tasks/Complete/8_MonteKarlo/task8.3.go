package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Task8_3 выполняет приближенное вычисление числа π методом Монте-Карло
func Task8_3(N int) {
	rand.Seed(time.Now().UnixNano())

	R := 7.0
	fmt.Printf("Задание 3\n")

	// 1. Выбираем количество случайных точек N
	// Оно передано как параметр функции Task8_3

	// 2. Генерация случайных точек и подсчет количества точек внутри круга
	M := countPointsInsideCircle(R, N)
	fmt.Printf("Количество попавших точек: %d\n", M)
	// 3. Приближенное вычисление числа π
	piApprox := 4.0 * float64(M) / float64(N)
	fmt.Printf("Приближенное значение числа π: %f\n", piApprox)

	// 4. Построение окружности и нанесение случайных точек на квадрат
	plotCircleWithPoints(R, N)
}

// Подсчет количества случайных точек, лежащих внутри круга радиуса R
func countPointsInsideCircle(R float64, N int) int {
	count := 0
	for i := 0; i < N; i++ {
		x := rand.Float64()*2*R - R
		y := rand.Float64()*2*R - R
		if x*x+y*y <= R*R {
			count++
		}
	}
	return count
}

// Построение окружности радиуса R и нанесение случайных точек на квадрат
func plotCircleWithPoints(R float64, N int) {
	p := plot.New()
	p.Title.Text = "Circle of Radius R and Random Points"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Min = -R
	p.X.Max = R
	p.Y.Min = -R
	p.Y.Max = R

	// Построение случайных точек
	randomPoints := make(plotter.XYs, N)
	for i := 0; i < N; i++ {
		x := rand.Float64()*2*R - R
		y := rand.Float64()*2*R - R
		randomPoints[i].X = x
		randomPoints[i].Y = y
	}
	points, err := plotter.NewScatter(randomPoints)
	if err != nil {
		panic(err)
	}
	points.GlyphStyle.Radius = vg.Points(1)
	p.Add(points)

	// Построение окружности радиуса R
	circlePoints := make(plotter.XYs, 361)
	for i := 0; i <= 360; i++ {
		angle := float64(i) * math.Pi / 180.0
		x := R * math.Cos(angle)
		y := R * math.Sin(angle)
		circlePoints[i].X = x
		circlePoints[i].Y = y
	}
	line, err := plotter.NewLine(circlePoints)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Width = vg.Points(1)
	p.Add(line)

	// Сохранение графика в файл
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "circle_with_points.png"); err != nil {
		panic(err)
	}
}
