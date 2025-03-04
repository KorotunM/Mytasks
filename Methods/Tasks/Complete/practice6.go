package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func round(value, precision float64) float64 {
	return math.Round(value/precision) * precision
}

func linearRegression(x, y []float64) (a, b float64) {
	n := float64(len(x))
	var sumX, sumY, sumXY, sumX2 float64

	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	b = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	a = (sumY - b*sumX) / n

	a = round(a, 0.01)
	b = round(b, 0.01)

	return a, b
}

func quadraticRegression(x, y []float64) (a, b, c float64) {
	n := float64(len(x))
	var sumX, sumX2, sumX3, sumX4, sumY, sumXY, sumX2Y float64

	for i := range x {
		sumX += x[i]
		sumX2 += x[i] * x[i]
		sumX3 += x[i] * x[i] * x[i]
		sumX4 += x[i] * x[i] * x[i] * x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2Y += x[i] * x[i] * y[i]
	}

	A := [3][3]float64{
		{n, sumX, sumX2},
		{sumX, sumX2, sumX3},
		{sumX2, sumX3, sumX4},
	}
	B := [3]float64{
		sumY,
		sumXY,
		sumX2Y,
	}

	// Решение системы линейных уравнений методом Крамера
	detA := A[0][0]*(A[1][1]*A[2][2]-A[1][2]*A[2][1]) -
		A[0][1]*(A[1][0]*A[2][2]-A[1][2]*A[2][0]) +
		A[0][2]*(A[1][0]*A[2][1]-A[1][1]*A[2][0])

	detA1 := B[0]*(A[1][1]*A[2][2]-A[1][2]*A[2][1]) -
		A[0][1]*(B[1]*A[2][2]-A[1][2]*B[2]) +
		A[0][2]*(B[1]*A[2][1]-A[1][1]*B[2])

	detA2 := A[0][0]*(B[1]*A[2][2]-A[1][2]*B[2]) -
		B[0]*(A[1][0]*A[2][2]-A[1][2]*A[2][0]) +
		A[0][2]*(A[1][0]*B[2]-B[1]*A[2][0])

	detA3 := A[0][0]*(A[1][1]*B[2]-B[1]*A[2][1]) -
		A[0][1]*(A[1][0]*B[2]-B[1]*A[2][0]) +
		B[0]*(A[1][0]*A[2][1]-A[1][1]*A[2][0])

	a = detA1 / detA
	b = detA2 / detA
	c = detA3 / detA

	a = round(a, 0.001)
	b = round(b, 0.001)
	c = round(c, 0.001)

	return a, b, c
}

func powerRegression(x, y []float64) (a, beta float64) {
	var logX, logY []float64
	for i := range x {
		logX = append(logX, math.Log(x[i]))
		logY = append(logY, math.Log(y[i]))
	}

	alpha, beta := linearRegression(logX, logY)
	a = math.Exp(alpha)

	a = round(a, 0.01)
	beta = round(beta, 0.01)

	return a, beta
}

func exponentialRegression(x, y []float64) (a, beta float64) {
	var logY []float64
	for i := range y {
		logY = append(logY, math.Log(y[i]))
	}

	alpha, beta := linearRegression(x, logY)
	a = math.Exp(alpha)

	a = round(a, 0.01)
	beta = round(beta, 0.01)

	return a, beta
}

func Task6() {
	//x := []float64{1, 2, 3, 4, 5, 6}
	//y := []float64{1.0, 1.5, 3.0, 4.5, 7.0, 8.5}

	x := []float64{3, 5, 7, 9, 11, 13}
	y := []float64{3.5, 4.4, 5.7, 6.1, 6.5, 7.3}

	a_lin, b_lin := linearRegression(x, y)
	fmt.Printf("Линейная функция: y = %.2fx + %.2f\n", b_lin, a_lin)

	yLin_pr := make([]float64, len(x))
	for i := range x {
		yLin_pr[i] = b_lin*x[i] + a_lin
	}

	errorLin := CalcError(y, yLin_pr)
	fmt.Printf("Суммарная погрешность для линейной аппроксимации: %.2f\n", errorLin)

	a_quad, b_quad, c_quad := quadraticRegression(x, y)
	fmt.Printf("Квадратичная функция: y = %.2fx^2 + %.2fx + %.2f\n", a_quad, b_quad, c_quad)

	yQuad_pr := make([]float64, len(x))
	for i := range x {
		yQuad_pr[i] = a_quad*x[i]*x[i] + b_quad*x[i] + c_quad
	}

	errorQuad := CalcError(y, yQuad_pr)
	fmt.Printf("Суммарная погрешность для квадратичной аппроксимации: %.2f\n", errorQuad)

	a_pow, beta_pow := powerRegression(x, y)
	fmt.Printf("Степенная функция: y = %.2fx^%.2f\n", a_pow, beta_pow)

	yPower_pr := make([]float64, len(x))
	for i := range x {
		yPower_pr[i] = a_pow * math.Pow(x[i], beta_pow)
	}

	errorPower := CalcError(y, yPower_pr)
	fmt.Printf("Суммарная погрешность для степенной аппроксимации: %.2f\n", errorPower)

	a_exp, beta_exp := exponentialRegression(x, y)
	fmt.Printf("Показательная функция: y = %.2fe^%.2fx\n", a_exp, beta_exp)

	yExp_pr := make([]float64, len(x))
	for i := range x {
		yExp_pr[i] = a_exp * math.Exp(beta_exp*x[i])
	}

	errorExp := CalcError(y, yExp_pr)
	fmt.Printf("Суммарная погрешность для показательной аппроксимации: %.2f\n", errorExp)

	// Подготовка данных для графика
	points := make(plotter.XYs, len(x))
	for i := range x {
		points[i].X = x[i]
		points[i].Y = y[i]
	}

	// Создание графика
	p := plot.New()
	p.Title.Text = "Аппроксимация методом наименьших квадратов"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Добавление экспериментальных точек
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Radius = vg.Points(2)
	p.Add(scatter)

	// Линейная функция
	line := plotter.NewFunction(func(x float64) float64 { return b_lin*x + a_lin })
	line.Color = plotutil.Color(1)

	// Квадратичная функция
	quadraticLine := plotter.NewFunction(func(x float64) float64 { return a_quad*x*x + b_quad*x + c_quad })
	quadraticLine.Color = plotutil.Color(2)

	// Степенная функция
	powerLine := plotter.NewFunction(func(x float64) float64 { return a_pow * math.Pow(x, beta_pow) })
	powerLine.Color = plotutil.Color(3)

	// Показательная функция
	expLine := plotter.NewFunction(func(x float64) float64 { return a_exp * math.Exp(beta_exp*x) })
	expLine.Color = plotutil.Color(4)

	// Настройка осей для лучшей визуализации
	p.Y.Min = 0  // Минимальное значение оси Y
	p.Y.Max = 10 // Максимальное значение оси Y
	p.X.Min = 0
	p.X.Max = 15
	// Изменение цвета линий
	line.Color = plotutil.Color(1)          // Линейная функция
	quadraticLine.Color = plotutil.Color(2) // Квадратичная функция
	powerLine.Color = plotutil.Color(3)     // Степенная функция
	expLine.Color = plotutil.Color(4)       // Показательная функция

	// Изменение типов линий (для различия)
	line.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	quadraticLine.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}

	// Добавление функций на график
	p.Add(line, quadraticLine, powerLine, expLine)
	p.Legend.Add("Линейная", line)
	p.Legend.Add("Квадратичная", quadraticLine)
	p.Legend.Add("Степенная", powerLine)
	p.Legend.Add("Показательная", expLine)
	p.Legend.Add("Точки", scatter)

	// Сохранение графика
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot_variant_7_extended.png"); err != nil {
		panic(err)
	}
}

func CalcError(y_now, y_func []float64) float64 {
	var sumError float64
	for i := range y_now {
		delta := y_now[i] - y_func[i]
		sumError += delta * delta
	}
	return round(sumError, 0.01)
}
