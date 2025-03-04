package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Compute one real root using Cardano's method
func solveCardanoOneRoot(a, b, c float64) float64 {
	// Reduced form: y^3 + py + q = 0
	p := b - (a * a / 3)
	q := c - (a*b)/3 + 2*(math.Pow(a/3, 3))

	// Compute s using the Cardano formula
	pDiv3 := p / 3
	qDiv2 := q / 2
	innerRoot := math.Sqrt(math.Pow(pDiv3, 3) + math.Pow(qDiv2, 2))
	s := math.Cbrt(-qDiv2 + innerRoot)
	t := math.Cbrt(-qDiv2 - innerRoot)

	// Compute y1 and x1
	y1 := s + t
	x1 := y1 - a/3
	return x1
}

// Solve quadratic equation with support for complex roots
func solveComplexQuadratic(a, b, c float64) (complex128, complex128) {
	discriminant := b*b - 4*a*c
	if discriminant >= 0 {
		// Real roots
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	} else {
		// Complex roots
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
		return complex(realPart, imaginaryPart), complex(realPart, -imaginaryPart)
	}
}

func Task10_Cardano() {
	// Values of alpha from small to very large
	alphaValues := []float64{1, 10, 112, 112345, 1e9, 1e12, 1e15, 1e16}

	fmt.Printf("\t=== Таблица корней и погрешностей для метода Кардано ===\n")
	fmt.Println("Alpha\t\t\tx1 \tError (x1)\t\tКомплексные x2 и x3\t\t\tОшибка корнях x2 и x3")

	for _, alpha := range alphaValues {
		// Define coefficients
		b := alpha * alpha
		c := 3 * alpha * alpha

		// Solve for the first root using Cardano's method
		a := 3.0
		x1 := solveCardanoOneRoot(a, b, c)
		temp := math.Pow(10, alpha)

		x1Theoretical := x1 - 1/temp
		x1Error := math.Abs(x1 - x1Theoretical)

		// Divide cubic equation by (x - x1) to get quadratic equation
		newB := 3 + x1
		newC := b + x1*x1 + 3*x1

		// Solve the quadratic equation for complex roots
		x2, x3 := solveComplexQuadratic(1, newB, newC)

		// Theoretical complex roots: -2, 0
		x2Theoretical := complex(0, -alpha)
		x3Theoretical := complex(0.0, -alpha)
		x2Error := cmplx.Abs(x2 - x2Theoretical)
		x3Error := cmplx.Abs(x3 - x3Theoretical)

		fmt.Printf("Alpha: %.2e\t  %.5f\t%.2e\t(%.2e, %.2ei) & (%.2e, %.2ei)\t%.2e & %.2e\n",
			alpha, x1, x1Error, real(x2), imag(x2), real(x3), imag(x3), x2Error, x3Error)
	}
}
