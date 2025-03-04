// explicit_scheme.go
package end

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func ExplicitScheme(dx, dt, T float64, nx, nt int) [][]float64 {
	alpha := dt / (dx * dx)
	u := make([][]float64, nt+1)
	for i := range u {
		u[i] = make([]float64, nx+1)
	}

	// Initial condition
	for i := 0; i <= nx; i++ {
		x := float64(i) * dx
		u[0][i] = math.Sin(x)
	}

	// Boundary conditions
	for n := 0; n <= nt; n++ {
		t := float64(n) * dt
		u[n][0] = 0
		u[n][nx] = math.Exp(t)
	}

	// Explicit scheme
	for n := 0; n < nt; n++ {
		for i := 1; i < nx; i++ {
			u[n+1][i] = u[n][i] + alpha*(u[n][i-1]-2*u[n][i]+u[n][i+1])
		}
	}

	return u
}

func PlotResults(u [][]float64, dx, dt float64, nx, nt int, filename string, title string) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "u"

	for n := 0; n <= nt; n += nt / 10 {
		pts := make(plotter.XYs, nx+1)
		for i := 0; i <= nx; i++ {
			x := float64(i) * dx
			pts[i].X = x
			pts[i].Y = u[n][i]
		}
		line, _ := plotter.NewLine(pts)
		line.LineStyle.Width = vg.Points(1)
		p.Add(line)
	}

	p.Save(6*vg.Inch, 4*vg.Inch, filename)
}

func Explicit() {
	L := math.Pi / 2
	T := 1.0
	nx := 50
	nt := 100
	dx := L / float64(nx)
	dt := T / float64(nt)

	uExplicit := ExplicitScheme(dx, dt, T, nx, nt)
	PlotResults(uExplicit, dx, dt, nx, nt, "explicit_scheme.png", "Explicit Scheme")
	fmt.Println("Explicit scheme results saved to explicit_scheme.png")
}
