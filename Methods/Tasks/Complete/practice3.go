package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Matrix struct {
	a float64
	b float64
	x float64
}

// Инициализация матрицы
func initialize(n int) [][]Matrix {
	var temp [][]Matrix
	if n == 2 {
		temp = [][]Matrix{
			{
				{a: 1e-4, b: 1, x: 0},
				{a: 1, b: 1, x: 0},
			},
			{
				{a: 1, b: 4, x: 0},
				{a: 2, b: 4, x: 0},
			},
		}
	} else if n == 3 {
		temp = [][]Matrix{
			{
				{a: 2.34, b: 14.41, x: 0},
				{a: -4.21, b: 0, x: 0},
				{a: -11.61, b: 0, x: 0},
			},
			{
				{a: 8.04, b: -6.44, x: 0},
				{a: 5.22, b: 0, x: 0},
				{a: 0.27, b: 0, x: 0},
			},
			{
				{a: 3.92, b: 55.56, x: 0},
				{a: -7.99, b: 0, x: 0},
				{a: 8.37, b: 0, x: 0},
			},
		}
	} else if n == 5 {
		temp = [][]Matrix{
			{
				{a: 4.43, b: 2.62, x: 0},
				{a: -7.21, b: 2.62, x: 0},
				{a: 8.05, b: 2.62, x: 0},
				{a: 1.23, b: 2.62, x: 0},
				{a: -2.56, b: 2.62, x: 0},
			},
			{
				{a: -1.29, b: -3.97, x: 0},
				{a: 6.47, b: -3.97, x: 0},
				{a: 2.96, b: -3.97, x: 0},
				{a: 3.22, b: -3.97, x: 0},
				{a: 6.12, b: -3.97, x: 0},
			},
			{
				{a: 6.12, b: -9.12, x: 0},
				{a: 8.31, b: -9.12, x: 0},
				{a: 9.41, b: -9.12, x: 0},
				{a: 1.78, b: -9.12, x: 0},
				{a: -2.88, b: -9.12, x: 0},
			},
			{
				{a: -2.57, b: 8.11, x: 0},
				{a: 6.93, b: 8.11, x: 0},
				{a: -3.74, b: 8.11, x: 0},
				{a: 7.41, b: 8.11, x: 0},
				{a: 5.55, b: 8.11, x: 0},
			},
			{
				{a: 1.46, b: 7.23, x: 0},
				{a: 3.62, b: 7.23, x: 0},
				{a: 7.83, b: 7.23, x: 0},
				{a: 6.25, b: 7.23, x: 0},
				{a: -2.35, b: 7.23, x: 0},
			},
		}

	}
	return temp
}

func swap_rows(mat *[][]Matrix, b []float64, n int, k int) {
	max := math.Abs((*mat)[k][k].a)
	maxRow := k
	for i := k + 1; i < n; i++ {
		if math.Abs((*mat)[i][k].a) > max {
			max = math.Abs((*mat)[i][k].a)
			maxRow = i
		}
	}
	if maxRow != k {
		(*mat)[k], (*mat)[maxRow] = (*mat)[maxRow], (*mat)[k]
		b[k], b[maxRow] = b[maxRow], b[k]
	}
}

func straight_run(mat *[][]Matrix, b []float64, n int) {
	for k := 0; k < n; k++ {
		swap_rows(mat, b, n, k)

		for i := k + 1; i < n; i++ {
			factor := (*mat)[i][k].a / (*mat)[k][k].a
			for j := k; j < n; j++ {
				(*mat)[i][j].a -= factor * (*mat)[k][j].a
			}
			b[i] -= factor * b[k]
		}
	}
}

func reverse_run(mat *[][]Matrix, b []float64, n int) []float64 {
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += (*mat)[i][j].a * x[j]
		}
		x[i] = (b[i] - sum) / (*mat)[i][i].a
	}
	return x
}

func calculate_residuals(mat *[][]Matrix, b []float64, x []float64, n int) []float64 {
	residuals := make([]float64, n)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += (*mat)[i][j].a * x[j]
		}
		residuals[i] = b[i] - sum
	}
	return residuals
}

func print_solution(out *bufio.Writer, x []float64, residuals []float64, n int) {
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "x%d = %f, Невязка: r = %e\n", i+1, x[i], residuals[i])
	}
}

// func Gaus(A *mat.Dense, b)
func Task3_3() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintf(out, "(A) Решения системы и их невязки: \n")
	mat := initialize(2)
	b := []float64{1, 4}
	straight_run(&mat, b, 2)
	x := reverse_run(&mat, b, 2)
	residuals := calculate_residuals(&mat, b, x, 2)
	print_solution(out, x, residuals, 2)

	fmt.Fprintf(out, "(Б) Решения системы и их невязки: \n")
	mat2 := initialize(3)
	b2 := []float64{14.41, -6.44, 55.56}
	straight_run(&mat2, b2, 3)
	x2 := reverse_run(&mat2, b2, 3)
	residuals2 := calculate_residuals(&mat2, b2, x2, 3)
	print_solution(out, x2, residuals2, 3)

	fmt.Fprintf(out, "(В) Решения системы и их невязки: \n")
	mat3 := initialize(5)
	b3 := []float64{2.62, -3.97, -9.12, 8.11, 7.23}
	straight_run(&mat3, b3, 5)
	x3 := reverse_run(&mat3, b3, 5)
	residuals3 := calculate_residuals(&mat3, b3, x3, 5)
	print_solution(out, x3, residuals3, 5)
}
