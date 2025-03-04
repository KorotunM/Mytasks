package end

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// analyticalSolution вычисляет аналитическое решение задачи
func analyticalSolution(x, t float64) float64 {
	return math.Exp(-t) * math.Sin(x)
}

// explicitScheme реализует явную разностную схему для решения задачи
func explicitScheme(L, T float64, Nx, Nt int) [][]float64 {
	dx := L / float64(Nx-1) // шаг по пространству
	dt := T / float64(Nt)   // шаг по времени
	r := dt / (dx * dx)     // параметр устойчивости

	if r > 0.5 {
		fmt.Printf("Предупреждение: Нарушено условие стабильности (r = %.2f > 0,5). Уменьшите значение dt или увеличьте значение Nx.\n", r)
	}

	u := make([][]float64, Nt) // сетка решения
	for i := range u {
		u[i] = make([]float64, Nx)
	}

	// Задаем начальные условия
	for i := 0; i < Nx; i++ {
		x := float64(i) * dx
		u[0][i] = math.Sin(x)
	}

	// Граничные условия
	for t := 0; t < Nt; t++ {
		time := float64(t) * dt
		u[t][0] = 0
		u[t][Nx-1] = math.Exp(-time)
	}

	// Явная схема
	for t := 0; t < Nt-1; t++ {
		for x := 1; x < Nx-1; x++ {
			u[t+1][x] = r*u[t][x-1] + (1-2*r)*u[t][x] + r*u[t][x+1]
		}
	}

	return u
}

func thomasAlgorithm(a, b, c, d []float64) []float64 {
	n := len(b)
	// Модифицированные коэффициенты
	cPrime := make([]float64, n)
	dPrime := make([]float64, n)

	// Прямая прогонка
	cPrime[0] = c[0] / b[0]
	dPrime[0] = d[0] / b[0]
	for i := 1; i < n; i++ {
		denominator := b[i] - a[i]*cPrime[i-1]
		cPrime[i] = c[i] / denominator
		dPrime[i] = (d[i] - a[i]*dPrime[i-1]) / denominator
	}

	// Обратная прогонка
	x := make([]float64, n)
	x[n-1] = dPrime[n-1]
	for i := n - 2; i >= 0; i-- {
		x[i] = dPrime[i] - cPrime[i]*x[i+1]
	}

	return x
}

// implicitScheme реализует неявную разностную схему для решения задачи
func implicitScheme(L, T float64, Nx, Nt int) [][]float64 {
	dx := L / float64(Nx-1) // шаг по пространству
	dt := T / float64(Nt)   // шаг по времени
	r := dt / (dx * dx)     // параметр устойчивости

	u := make([][]float64, Nt) // сетка решения
	for i := range u {
		u[i] = make([]float64, Nx)
	}

	// Задаем начальные условия
	for i := 0; i < Nx; i++ {
		x := float64(i) * dx
		u[0][i] = math.Sin(x)
	}

	// Граничные условия
	for t := 0; t < Nt; t++ {
		time := float64(t) * dt
		u[t][0] = 0
		u[t][Nx-1] = math.Exp(-time)
	}

	// Коэффициенты трёхдиагональной матрицы
	a := make([]float64, Nx-2)
	b := make([]float64, Nx-2)
	c := make([]float64, Nx-2)
	d := make([]float64, Nx-2)

	for i := 0; i < Nx-2; i++ {
		a[i] = -r
		b[i] = 1 + 2*r
		c[i] = -r
	}

	// Неявная схема с использованием метода Томаса
	for t := 0; t < Nt-1; t++ {
		// Формируем правую часть системы
		for i := 0; i < Nx-2; i++ {
			d[i] = u[t][i+1]
		}
		d[0] += r * u[t+1][0]
		d[Nx-3] += r * u[t+1][Nx-1]

		// Решаем систему методом Томаса
		solution := thomasAlgorithm(a, b, c, d)

		// Записываем результат в сетку
		for i := 0; i < Nx-2; i++ {
			u[t+1][i+1] = solution[i]
		}
	}

	return u
}

// calculateError вычисляет среднюю ошибку между численным и аналитическим решением
func calculateError(u [][]float64, L, T float64, Nx, Nt int) float64 {
	dx := L / float64(Nx-1)
	dt := T / float64(Nt)
	totalError := 0.0
	count := 0

	for t := 0; t < Nt; t++ {
		time := float64(t) * dt
		for i := 0; i < Nx; i++ {
			x := float64(i) * dx
			analytical := analyticalSolution(x, time)
			error := math.Abs(u[t][i] - analytical)
			totalError += error
			count++
		}
	}

	return totalError / float64(count)
}

// plotResults2 визуализирует результаты схем
func plotResults2(u [][]float64, L, T float64, Nx, Nt int, filename, title string) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "u"

	dx := L / float64(Nx-1)                                   // шаг по пространству
	timesteps := []int{0, Nt / 4, Nt / 2, 3 * Nt / 4, Nt - 1} // временные слои для визуализации
	for _, t := range timesteps {
		linePts := make(plotter.XYs, Nx)
		nodePts := make(plotter.XYs, Nx)
		for i := 0; i < Nx; i++ {
			x := float64(i) * dx
			linePts[i].X = x
			linePts[i].Y = u[t][i]
			nodePts[i].X = x
			nodePts[i].Y = u[t][i]
		}
		line, _ := plotter.NewLine(linePts)
		line.LineStyle.Width = vg.Points(1)
		p.Add(line)

		points, _ := plotter.NewScatter(nodePts)
		points.GlyphStyle.Radius = vg.Points(2)
		p.Add(points)
	}

	p.Save(6*vg.Inch, 4*vg.Inch, filename)
}
func plotAnalyticalSolution(L, T float64, Nx, Nt int, filename string) {
	p := plot.New()
	p.Title.Text = "Analytical Solution"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "u"

	dx := L / float64(Nx-1)
	dt := T / float64(Nt)

	for t := 0; t <= Nt; t += Nt / 5 {
		pts := make(plotter.XYs, Nx)
		for i := 0; i < Nx; i++ {
			x := float64(i) * dx
			time := float64(t) * dt
			pts[i].X = x
			pts[i].Y = analyticalSolution(x, time)
		}
		line, _ := plotter.NewLine(pts)
		line.LineStyle.Width = vg.Points(1)
		p.Add(line)
	}

	p.Save(6*vg.Inch, 4*vg.Inch, filename)
}
func PyToGo() {
	L := math.Pi / 2  // длина области
	T := 3.0          // время моделирования
	Nx := 10          // количество узлов по пространству
	Nt := 2 * Nx * Nx // количество шагов по времени (условие стабильности)

	uExplicit := explicitScheme(L, T, Nx, Nt)
	plotResults2(uExplicit, L, T, Nx, Nt, "explicit_scheme_with_nodes.png", "Explicit Scheme with Nodes")
	fmt.Println("Явное решение сохранено как explicit_scheme_with_nodes.png")

	uImplicit := implicitScheme(L, T, Nx, Nt)
	plotResults2(uImplicit, L, T, Nx, Nt, "implicit_scheme_with_nodes.png", "Implicit Scheme with Nodes")
	fmt.Println("Неявное решение сохранено как implicit_scheme_with_nodes.png")

	plotAnalyticalSolution(L, T, Nx, Nt, "analytical_solution.png")
	fmt.Println("Аналитическое решение сохранено как analytical_solution.png")

	errorExplicit := calculateError(uExplicit, L, T, Nx, Nt)
	fmt.Printf("Среднняя ошибка для Явной: %f\n", errorExplicit)

	errorImplicit := calculateError(uImplicit, L, T, Nx, Nt)
	fmt.Printf("Среднняя ошибка для Неявной: %f\n", errorImplicit)
}
