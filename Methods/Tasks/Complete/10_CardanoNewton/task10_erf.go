package main

import (
	"fmt"
	"math"
	"tasks/tasks/pkg"
)

// Define the equation and its derivative
func erfEquation(x float64) float64 {
	return pkg.Erf(x) - 0.5
}

func erfDerivative(x float64) float64 {
	// Derivative of erf(x) is (2/sqrt(pi)) * exp(-x^2)
	return (2 / math.Sqrt(math.Pi)) * math.Exp(-x*x)
}

// Newton's Method
func newtonErf(tolerance float64, maxIterations int) (float64, int, error) {
	// Initial guess
	x := 1.2 // Starting point near the solution

	for i := 0; i < maxIterations; i++ {
		fx := erfEquation(x)
		fpx := erfDerivative(x)

		if math.Abs(fpx) < 1e-15 { // Avoid division by zero
			return x, i, fmt.Errorf("derivative is too small at iteration %d", i)
		}

		xNext := x - fx/fpx

		if math.Abs(xNext-x) < tolerance {
			return xNext, i + 1, nil
		}

		x = xNext
	}

	return x, maxIterations, fmt.Errorf("method did not converge within the maximum number of iterations")
}

func Task10_erf() {
	// Parameters for Newton's method
	tolerance := 1e-9
	maxIterations := 100

	fmt.Println("=== Solving erf(x) - 0.5 = 0 using Newton's Method ===")

	// Solve the equation
	root, iterations, err := newtonErf(tolerance, maxIterations)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Root: %.13f found in %d iterations\n", root, iterations)
	}
}
