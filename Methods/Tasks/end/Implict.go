// explicit_scheme.go
package end

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func ImplicitScheme(dx, dt, T float64, nx, nt int) [][]float64 {
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

	// Implicit scheme (Thomas algorithm for tridiagonal matrix)
	A := make([]float64, nx-1)
	B := make([]float64, nx-1)
	C := make([]float64, nx-1)
	D := make([]float64, nx-1)

	for n := 0; n < nt; n++ {
		for i := 0; i < nx-1; i++ {
			A[i] = -alpha
			B[i] = 1 + 2*alpha
			C[i] = -alpha
			D[i] = u[n][i+1]
		}
		D[0] += alpha * u[n+1][0]
		D[nx-2] += alpha * u[n+1][nx]

		// Forward sweep
		for i := 1; i < nx-1; i++ {
			m := A[i] / B[i-1]
			B[i] -= m * C[i-1]
			D[i] -= m * D[i-1]
		}

		// Back substitution
		u[n+1][nx-1] = D[nx-2] / B[nx-2]
		for i := nx - 3; i >= 0; i-- {
			u[n+1][i+1] = (D[i] - C[i]*u[n+1][i+2]) / B[i]
		}
	}

	return u
}

func plotResults(u [][]float64, dx, dt float64, nx, nt int, filename string, title string) {
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

func Implicit() {
	L := math.Pi / 2
	T := 1.0
	nx := 50
	nt := 100
	dx := L / float64(nx)
	dt := T / float64(nt)

	uImplicit := ImplicitScheme(dx, dt, T, nx, nt)
	plotResults(uImplicit, dx, dt, nx, nt, "implicit_scheme.png", "Implicit Scheme")
	fmt.Println("Implicit scheme results saved to implicit_scheme.png")
}
