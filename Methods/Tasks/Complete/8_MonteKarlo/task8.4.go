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

// Task8_4 выполняет приближенное вычисление площади фигуры методом Монте-Карло
// Фигура задана замкнутой кривой в полярных координатах A*cos^2(φ) + B*sin^2(φ) = ρ^2
func Task8_4(N int) {
	rand.Seed(time.Now().UnixNano())

	A := 18.0
	B := 4.0
	fmt.Printf("Задание 4\n")

	// 1. Построение кривой и переход от полярных координат к декартовым
	plotClosedCurve(A, B)

	// 2. Определение размеров прямоугольника, в котором лежит фигура
	a := math.Sqrt(A)
	b := math.Sqrt(A)

	// 3. Генерация случайных точек и подсчет количества точек внутри фигуры
	M := countPointsInsideCurve(A, B, a, b, N)
	fmt.Printf("Количество случайных точек %d\n", M)
	// 4. Приближенное вычисление площади фигуры
	areaApprox := 4.0 * a * b * float64(M) / float64(N)
	fmt.Printf("Приближенная площадь фигуры: %f\n", areaApprox)
}

// Построение кривой и переход от полярных координат к декартовым
func plotClosedCurve(A, B float64) {
	p := plot.New()
	p.Title.Text = "Closed Curve Defined by A*cos^2(φ) + B*sin^2(φ) = ρ^2"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Min = -math.Sqrt(A)
	p.X.Max = math.Sqrt(A)
	p.Y.Min = -math.Sqrt(A)
	p.Y.Max = math.Sqrt(A)

	// Построение кривой
	n := 361
	curvePoints := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		phi := float64(i) * 2 * math.Pi / float64(n-1)
		rho := math.Sqrt(A*math.Pow(math.Cos(phi), 2) + B*math.Pow(math.Sin(phi), 2))
		x := rho * math.Cos(phi)
		y := rho * math.Sin(phi)
		curvePoints[i].X = x
		curvePoints[i].Y = y
	}

	line, err := plotter.NewLine(curvePoints)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Width = vg.Points(1)
	p.Add(line)

	// Сохранение графика в файл
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "closed_curve.png"); err != nil {
		panic(err)
	}
}

// Подсчет количества случайных точек, лежащих внутри фигуры, заданной кривой
func countPointsInsideCurve(A, B, a, b float64, N int) int {
	count := 0
	for i := 0; i < N; i++ {
		x := rand.Float64()*2*a - a
		y := rand.Float64()*2*b - b
		rho := math.Sqrt(x*x + y*y)
		phi := math.Atan2(y, x)
		if phi < 0 {
			phi += 2 * math.Pi
		}
		rhoCurve := math.Sqrt(A*math.Pow(math.Cos(phi), 2) + B*math.Pow(math.Sin(phi), 2))
		if rho <= rhoCurve {
			count++
		}
	}
	return count
}
