package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Определим функцию, которую будем интерполировать
func OrigF(x float64) float64 {
	return 1 / (1 + 25*x*x)
}

// Генерация равноотстоящих узлов
func generateEquidistantNodes(n int, a, b float64) []float64 {
	nodes := make([]float64, n)
	for i := 0; i < n; i++ {
		nodes[i] = a + (b-a)*float64(i)/float64(n-1)
	}
	return nodes
}

// Построение кубического сплайна для заданной таблицы значений
func cubicSplineInterpolation(x, xVals, yVals []float64) []float64 {
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
		dx := v - xVals[segment]
		fmt.Printf("a = %f, b = %f, c=%f,d=%f\n", a[segment], b[segment], c[segment], d[segment])
		result[i] = a[segment] + b[segment]*dx + c[segment]*dx*dx + d[segment]*dx*dx*dx
	}

	return result
}

// Функция для построения графиков
func generatePlotGraphs(n int, a, b float64) {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Function and Cubic Spline Interpolation (n=%d)", n)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Y.Min = -0.5
	p.Y.Max = 1.5

	// Исходная функция
	fLine := plotter.NewFunction(OrigF)
	fLine.Color = plotter.DefaultLineStyle.Color
	fLine.Width = vg.Points(2)
	p.Add(fLine)
	p.Legend.Add("f(x)", fLine)

	// Построение графика для равноотстоящих узлов
	equidistantNodes := generateEquidistantNodes(n, a, b)
	equidistantValues := make([]float64, len(equidistantNodes))
	for i, x := range equidistantNodes {
		equidistantValues[i] = f(x)
	}

	// Генерация точек для построения графика сплайна
	xPoints := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		xPoints[i] = a + (b-a)*float64(i)/999
	}
	ySpline := cubicSplineInterpolation(xPoints, equidistantNodes, equidistantValues)

	splinePoints := make(plotter.XYs, len(xPoints))
	for i := range xPoints {
		splinePoints[i].X = xPoints[i]
		splinePoints[i].Y = ySpline[i]
	}

	splineLine, err := plotter.NewLine(splinePoints)
	if err != nil {
		panic(err)
	}
	splineLine.Color = plotter.DefaultLineStyle.Color
	splineLine.Width = vg.Points(1)
	splineLine.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	p.Add(splineLine)
	p.Legend.Add("Cubic Spline", splineLine)

	// Сохранение графика
	filename := fmt.Sprintf("cubic_spline_n%d.png", n)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, filename); err != nil {
		panic(err)
	}
	fmt.Printf("Saved graph to %s\n", filename)
}

func Part_2() {
	a := -1.0
	b := 1.0

	// Построение графиков для n = 5, 10, 15
	for _, n := range []int{5, 10, 15} {
		generatePlotGraphs(n, a, b)
	}
	//пункт 2
	//generatePlotGraphsAndCheckSpline()
}
