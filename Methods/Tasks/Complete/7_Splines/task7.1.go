package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Определим функцию, которую будем интерполировать
func f(x float64) float64 {
	return 1 / (1 + 25*x*x)
}

// Генерация равноотстоящих узлов
func equidistantNodes(n int, a, b float64) []float64 {
	nodes := make([]float64, n)
	for i := 0; i < n; i++ {
		nodes[i] = a + (b-a)*float64(i)/float64(n-1)
	}
	return nodes
}

// Генерация узлов Чебышева
func chebyshevNodes(n int, a, b float64) []float64 {
	nodes := make([]float64, n)
	for i := 0; i < n; i++ {
		nodes[i] = 0.5 * ((b-a)*math.Cos(math.Pi*(2*float64(i)+1)/(2*float64(n))) + (a + b))
	}
	return nodes
}

// Интерполяционный полином Лагранжа
func lagrangeInterpolation(x float64, nodes []float64, values []float64) float64 {
	n := len(nodes)
	result := 0.0
	for i := 0; i < n; i++ {
		term := values[i]
		for j := 0; j < n; j++ {
			if i != j {
				term *= (x - nodes[j]) / (nodes[i] - nodes[j])
			}
		}
		result += term
	}
	return result
}

// Функция для построения графиков
func plotGraphs(n int, a, b float64) {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Интерполяция для (n=%d)", n)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Y.Min = -0.5
	p.Y.Max = 1.5

	// Исходная функция
	fLine := plotter.NewFunction(f)
	fLine.Color = plotter.DefaultLineStyle.Color
	fLine.Width = vg.Points(2)
	p.Add(fLine)
	p.Legend.Add("f(x)", fLine)

	// Построение графика для равноотстоящих узлов
	equidistantNodes := equidistantNodes(n, a, b)
	equidistantValues := make([]float64, len(equidistantNodes))
	for i, x := range equidistantNodes {
		equidistantValues[i] = f(x)
	}

	equidistantPoints := make(plotter.XYs, 1000)
	for i := range equidistantPoints {
		x := a + (b-a)*float64(i)/999
		equidistantPoints[i].X = x
		equidistantPoints[i].Y = lagrangeInterpolation(x, equidistantNodes, equidistantValues)
	}
	equidistantLine, err := plotter.NewLine(equidistantPoints)
	if err != nil {
		panic(err)
	}
	equidistantLine.Color = plotter.DefaultLineStyle.Color
	equidistantLine.Width = vg.Points(1)
	equidistantLine.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	p.Add(equidistantLine)
	p.Legend.Add("с равноотстоящими узлами", equidistantLine)

	// Построение графика для узлов Чебышева
	chebyshevNodes := chebyshevNodes(n, a, b)
	chebyshevValues := make([]float64, len(chebyshevNodes))
	for i, x := range chebyshevNodes {
		chebyshevValues[i] = f(x)
	}

	chebyshevPoints := make(plotter.XYs, 1000)
	for i := range chebyshevPoints {
		x := a + (b-a)*float64(i)/999
		chebyshevPoints[i].X = x
		chebyshevPoints[i].Y = lagrangeInterpolation(x, chebyshevNodes, chebyshevValues)
	}
	chebyshevLine, err := plotter.NewLine(chebyshevPoints)
	if err != nil {
		panic(err)
	}
	chebyshevLine.Color = plotter.DefaultLineStyle.Color
	chebyshevLine.Width = vg.Points(1)
	chebyshevLine.Dashes = []vg.Length{vg.Points(2), vg.Points(3)}
	p.Add(chebyshevLine)
	p.Legend.Add("с чебышевскими узлами", chebyshevLine)

	// Сохранение графика
	filename := fmt.Sprintf("lagrange_combined_n%d.png", n)
	if err := p.Save(8*vg.Inch, 8*vg.Inch, filename); err != nil {
		panic(err)
	}
	fmt.Printf("Saved graph to %s\n", filename)
}

func main() {
	a := -1.0
	b := 1.0

	// Построение графиков для n = 5, 10, 15 с равноотстоящими узлами и узлами Чебышева на одном графике
	for _, n := range []int{5, 10, 15} {
		plotGraphs(n, a, b)
	}
	Part_2()
}
