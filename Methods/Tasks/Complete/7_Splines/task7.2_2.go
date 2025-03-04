package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Построение кубического сплайна для заданной таблицы значений
func cubicSplineInterpolation_2(x, xVals, yVals []float64) []float64 {
	n := len(xVals)
	if n < 3 {
		panic("Need at least three points for cubic spline interpolation")
	}

	h := make([]float64, n-1)
	alpha := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		h[i] = xVals[i+1] - xVals[i]
		if i > 0 {
			alpha[i] = (3/h[i])*(yVals[i+1]-yVals[i]) - (3/h[i-1])*(yVals[i]-yVals[i-1])
		}
	}

	l := make([]float64, n)
	mu := make([]float64, n)
	z := make([]float64, n)
	l[0] = 1
	mu[0] = 0
	z[0] = 0

	for i := 1; i < n-1; i++ {
		l[i] = 2*(xVals[i+1]-xVals[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}

	l[n-1] = 1
	z[n-1] = 0
	c := make([]float64, n)
	b := make([]float64, n-1)
	d := make([]float64, n-1)
	a := make([]float64, n-1)

	for j := n - 2; j >= 0; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
		b[j] = (yVals[j+1]-yVals[j])/h[j] - h[j]*(c[j+1]+2*c[j])/3
		d[j] = (c[j+1] - c[j]) / (3 * h[j])
		a[j] = yVals[j]
	}

	// Генерация точек для построения графика сплайна
	result := make([]float64, len(x))
	for i, v := range x {
		// Найти нужный сегмент
		segment := 0
		for j := 0; j < n-1; j++ {
			if v >= xVals[j] && v <= xVals[j+1] {
				segment = j
				break
			}
		}
		fmt.Printf("a = %f, b = %f, c=%f,d=%f\n", a[segment], b[segment], c[segment], d[segment])
		dx := v - xVals[segment]
		result[i] = a[segment] + b[segment]*dx + c[segment]*dx*dx + d[segment]*dx*dx*dx

	}

	return result
}

// Функция для построения графиков и проверки интерполяции
func generatePlotGraphsAndCheckSpline() {
	p := plot.New()
	p.Title.Text = "Cubic Spline for Given Table"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Заданные значения таблицы
	xVals := []float64{2, 3, 5, 7}
	yVals := []float64{4, -2, 6, -3}

	// Генерация точек для построения графика сплайна
	xSpline := make([]float64, 100)
	for i := 0; i < 100; i++ {
		xSpline[i] = 2 + 5*float64(i)/99
	}
	ySpline := cubicSplineInterpolation_2(xSpline, xVals, yVals)

	splinePoints := make(plotter.XYs, len(xSpline))
	for i := range xSpline {
		splinePoints[i].X = xSpline[i]
		splinePoints[i].Y = ySpline[i]
	}

	splineLine, err := plotter.NewLine(splinePoints)
	if err != nil {
		panic(err)
	}
	splineLine.Color = plotter.DefaultLineStyle.Color
	splineLine.Width = vg.Points(2)
	p.Add(splineLine)
	p.Legend.Add("Cubic Spline")

	// Добавление узловых точек на график
	points := make(plotter.XYs, len(xVals))
	for i := range xVals {
		points[i].X = xVals[i]
		points[i].Y = yVals[i]
	}
	nodePoints, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	nodePoints.GlyphStyle.Radius = vg.Points(3)
	nodePoints.GlyphStyle.Color = plotter.DefaultLineStyle.Color
	p.Add(nodePoints)
	p.Legend.Add("Nodes")

	// Сохранение графика
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "cubic_spline_table.png"); err != nil {
		panic(err)
	}
	fmt.Printf("Saved graph to cubic_spline_table.png\n")

	// Проверка интерполяции в узловых точках
	fmt.Println("Interpolation results at node points:")
	for i, x := range xVals {
		yInterpolated := cubicSplineInterpolation_2([]float64{x}, xVals, yVals)[0]
		fmt.Printf("x = %v, y_actual = %v, y_interpolated = %v\n", x, yVals[i], yInterpolated)
	}
}
