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

// Task8_2 выполняет приближенное вычисление интеграла по методу Монте-Карло
func Part_2(N int) {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Задание 2\n")

	// 1. Построение графика подынтегральной функции на интервале [0, 5]
	plotSubIntegralGraph()

	// 2. Выбираем количество случайных точек N
	// Оно передано как параметр функции Part_2

	// 3. Определение прямоугольника, в котором лежит график
	a := 5.0
	b := calculateMaxY()

	// 4. Генерация случайных точек и подсчет количества точек внутри фигуры
	M := countPointsInsideIntegral(a, b, N)
	fmt.Printf("Количество попавших точек: %d\n", M)
	// 5. Вычисление площади по формуле Монте-Карло
	areaApprox := approximateArea_2(a, b, M, N)
	fmt.Printf("Приближенная площадь фигуры: %f\n", areaApprox)

	// 6. Оценка абсолютной и относительной погрешности
	trueArea := calculateTrueAreaIntegral()
	absoluteError := math.Abs(trueArea - areaApprox)
	relativeError := absoluteError / trueArea

	fmt.Printf("Абсолютная погрешность: %f\n", absoluteError)
	fmt.Printf("Относительная погрешность: %f\n", relativeError)
}

// Построение графика подынтегральной функции на интервале [0, 5]
func plotSubIntegralGraph() {
	p := plot.New()
	p.Title.Text = "Graph of f(x) = sqrt(11 - 7*sin^2(x))"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Min = 0
	p.X.Max = 5
	p.Y.Min = 0
	p.Y.Max = calculateMaxY()

	// Основной график функции f(x) = sqrt(11 - 7*sin^2(x))
	n := 501
	xys := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		x := float64(i) * 0.01
		y := math.Sqrt(11 - 7*math.Pow(math.Sin(x), 2))
		xys[i].X = x
		xys[i].Y = y
	}

	line, err := plotter.NewLine(xys)
	if err != nil {
		panic(err)
	}
	line.Color = plotter.DefaultLineStyle.Color
	p.Add(line)

	// Построение линий x = 0, x = 5, y = 0
	verticalLineX0 := make(plotter.XYs, 2)
	verticalLineX0[0].X = 0
	verticalLineX0[0].Y = 0
	verticalLineX0[1].X = 0
	verticalLineX0[1].Y = calculateMaxY()

	verticalLineX5 := make(plotter.XYs, 2)
	verticalLineX5[0].X = 5
	verticalLineX5[0].Y = 0
	verticalLineX5[1].X = 5
	verticalLineX5[1].Y = calculateMaxY()

	horizontalLineY0 := make(plotter.XYs, 2)
	horizontalLineY0[0].X = 0
	horizontalLineY0[0].Y = 0
	horizontalLineY0[1].X = 5
	horizontalLineY0[1].Y = 0

	lineX0, err := plotter.NewLine(verticalLineX0)
	if err != nil {
		panic(err)
	}
	lineX5, err := plotter.NewLine(verticalLineX5)
	if err != nil {
		panic(err)
	}
	lineY0, err := plotter.NewLine(horizontalLineY0)
	if err != nil {
		panic(err)
	}

	lineX0.LineStyle.Width = vg.Points(1)
	lineX5.LineStyle.Width = vg.Points(1)
	lineY0.LineStyle.Width = vg.Points(1)

	p.Add(lineX0, lineX5, lineY0)

	// Сохранение графика в файл
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "sub_integral_graph.png"); err != nil {
		panic(err)
	}
}

// Функция для подсчета количества случайных точек, лежащих внутри интегральной фигуры
func countPointsInsideIntegral(a, b float64, N int) int {
	count := 0
	for i := 0; i < N; i++ {
		x := rand.Float64() * a
		y := rand.Float64() * b
		if y < math.Sqrt(11-7*math.Pow(math.Sin(x), 2)) {
			count++
		}
	}
	return count
}

// Приближенное вычисление площади по методу Монте-Карло
func approximateArea_2(a, b float64, M int, N int) float64 {
	return float64(M) / float64(N) * a * b
}

// Функция для вычисления максимального значения Y на графике интегральной функции
func calculateMaxY() float64 {
	return math.Sqrt(11)
}

// Вычисление истинной площади фигуры для интеграла (аналитическое решение для проверки погрешности)
func calculateTrueAreaIntegral() float64 {
	// Примерное значение истинной площади интеграла
	return 13.4 // Поставьте соответствующее значение, если доступно аналитическое решение
}
