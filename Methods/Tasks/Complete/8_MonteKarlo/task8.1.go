package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	n int = 7   // пример значения n (необходимо подставить ваш вариант)
	N int = 100 // количество случайных точек
)

func Part_1() {
	rand.Seed(time.Now().UnixNano())

	if n <= 10 {
		// Построение графика для функции f(x)
		plotGraphF()
	} else {
		// Построение графиков для функций f1(x) и f2(x)
		plotGraphF1F2()
	}

	// Определение прямоугольника и подсчет случайных точек
	a, b := determineRectangleDimensions()
	fmt.Printf("Задание 1 \n")
	M := countPointsInside(a, b)
	fmt.Printf("Количество случайных точек: %d\n", M)
	// Вычисление площади по формуле Монте-Карло
	areaApprox := approximateArea(a, b, M)
	fmt.Printf("Приближенная площадь: %f\n", areaApprox)

	// Вычисление абсолютной и относительной погрешности
	trueArea := calculateTrueArea()
	absoluteError := math.Abs(trueArea - areaApprox)
	relativeError := absoluteError / trueArea

	fmt.Printf("Абсолютная погрешность: %f\n", absoluteError)
	fmt.Printf("Относительная погрешность: %f\n", relativeError)
}

// Построение графика функции f(x) для n <= 10
func plotGraphF() {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Graph of f(x)"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "X", Min: 0, Max: 20}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Y", Min: 0, Max: 20}),
	)

	xData := make([]float64, 21)
	yData := make([]float64, 21)
	for i := 0; i <= 20; i++ {
		xData[i] = float64(i)
		if float64(i) >= 0 && float64(i) <= float64(n) {
			yData[i] = (10 * float64(i)) / float64(n)
		} else if float64(i) > float64(n) && float64(i) <= 20 {
			yData[i] = 10 * (float64(i) - 20) / float64(n-20)
		} else {
			yData[i] = 0
		}
	}

	line.SetXAxis(xData).
		AddSeries("f(x)", generateLineItems(xData, yData)).
		AddSeries("x = 0", generateLineItems([]float64{0, 0}, []float64{0, 20})).
		AddSeries("y = 0", generateLineItems([]float64{0, 20}, make([]float64, 21)))

	f, err := os.Create("f_graph.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	line.Render(f)
}

// Построение графиков функций f1(x) и f2(x) для n > 10
func plotGraphF1F2() {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Graphs of f1(x) and f2(x)"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "X", Min: 0, Max: 30}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Y", Min: 0, Max: 70}),
	)

	xData := make([]float64, 31)
	f1Data := make([]float64, 31)
	f2Data := make([]float64, 31)
	for i := 0; i <= 30; i++ {
		xData[i] = float64(i)
		f1Data[i] = (10 * float64(i)) / float64(n)
		f2Data[i] = 10*(float64(i)-20)/float64(n-20) + 20
	}

	line.SetXAxis(xData).
		AddSeries("f1(x)", generateLineItems(xData, f1Data)).
		AddSeries("f2(x)", generateLineItems(xData, f2Data)).
		AddSeries("x = 0", generateLineItems([]float64{0, 0}, []float64{0, 70}))

	f, err := os.Create("f1_f2_graph.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	line.Render(f)
}

// Функция для генерации данных для линии
func generateLineItems(xData, yData []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(xData); i++ {
		if !math.IsNaN(yData[i]) {
			items = append(items, opts.LineData{Value: yData[i]})
		}
	}
	return items
}

// Подсчет количества случайных точек, лежащих внутри фигуры
func countPointsInside(a, b float64) int {
	count := 0
	for i := 0; i < N; i++ {
		x := rand.Float64() * a
		y := rand.Float64() * b
		if n <= 10 {
			if y < f(x) {
				count++
			}
		} else {
			if x >= 0 && x <= float64(n) {
				if y < f1(x) {
					count++
				}
			} else if x > float64(n) && x <= 20 {
				if y < f2(x) {
					count++
				}
			}
		}
	}
	return count
}

// Приближенное вычисление площади по методу Монте-Карло
func approximateArea(a, b float64, M int) float64 {
	return float64(M) / float64(N) * a * b
}

// Функция f(x) для n <= 10
func f(x float64) float64 {
	if x >= 0 && x <= float64(n) {
		return (10 * x) / float64(n)
	} else if x > float64(n) && x <= 20 {
		return 10 * (x - 20) / float64(n-20)
	}
	return 0
}

// Функция f1(x) для n > 10
func f1(x float64) float64 {
	if x >= 0 && x <= float64(n) {
		return (10 * x) / float64(n)
	}
	return math.NaN()
}

// Функция f2(x) для n > 10
func f2(x float64) float64 {
	if x > float64(n) && x <= 20 {
		return 10 * (x - 20) / float64(n-20)
	}
	return math.NaN()
}

// Вычисление истинной площади фигуры для вычисления погрешности
func calculateTrueArea() float64 {
	if n <= 10 {
		// Вычисление площади треугольника для f(x) при n <= 10
		return 100
	} else {
		// Вычисление площади фигуры для f1(x) и f2(x) при n > 10
		area1 := 0.5 * float64(n) * (10 * float64(n) / float64(n))
		area2 := 0.5 * (20 - float64(n)) * (10 * (20 - float64(n)) / float64(n-20))
		return area1 + area2
	}
}
func determineRectangleDimensions() (float64, float64) {
	var a, b float64
	if n <= 10 {
		a = 20.0
		b = 10.0
	} else {
		a = 20.0
		b = 70.0 // для n > 10 максимум по y будет достигать 70
	}
	return a, b
}
