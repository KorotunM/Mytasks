package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"

// 	"gonum.org/v1/gonum/mat"
// 	//"gonum.org/v1/gonum/stat/distuv"
// )

// // Функция вычисления Erf(x) из библиотеки Gonum
// func NewA(flag int) *mat.Dense {
// 	var A *mat.Dense
// 	if flag == 1 {
// 		A = mat.NewDense(3, 3, []float64{
// 			1.00, 0.80, 0.64,
// 			1.00, 0.90, 0.81,
// 			1.00, 1.10, 1.21,
// 		})
// 	} else if flag == 2 {
// 		A = mat.NewDense(3, 3, []float64{
// 			0.1, 0.2, 0.3,
// 			0.4, 0.5, 0.6,
// 			0.7, 0.8, 0.9,
// 		})
// 	}
// 	return A
// }
// func Newb(flag int) *mat.VecDense {
// 	var b *mat.VecDense
// 	if flag == 1 {
// 		b = mat.NewVecDense(3, []float64{
// 			Erf(0.80),
// 			Erf(0.90),
// 			Erf(1.10),
// 		})
// 	} else if flag == 2 {
// 		b = mat.NewVecDense(3, []float64{
// 			0.1, 0.3, 0.5,
// 		})
// 	}
// 	return b
// }
// func Task4_4() {
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	fmt.Fprintln(out, "(A) Решение системы: ")
// 	A := NewA(1)
// 	b := Newb(1)
// 	var x mat.VecDense
// 	newA, newB := Gaus(A, b)
// 	err := x.SolveVec(newA, newB)
// 	if err != nil {
// 		fmt.Fprint(out, "Ошибка решения системы:", err)
// 		return
// 	}

// 	fmt.Fprintf(out, "Решение системы:\n x1 = %.6f\n x2 = %.6f\n x3 = %.6f\n", x.AtVec(0), x.AtVec(1), x.AtVec(2))

// 	sumX := x.AtVec(0) + x.AtVec(1) + x.AtVec(2)
// 	fmt.Fprintf(out, "Сумма решений x1 + x2 + x3 = %.6f\n", sumX)

// 	Erf1 := Erf(1.0)
// 	fmt.Fprintf(out, "Значение Erf(1.0) = %.6f\n", Erf1)
// 	fmt.Fprintf(out, "Разница = %.6f\n", math.Abs(sumX-Erf1))

// 	fmt.Fprintln(out, "(B) Решение системы: ")
// 	A2 := NewA(2)
// 	det := mat.Det(A2)
// 	if det == 0 {
// 		fmt.Fprintln(out, "Матрица вырожденная, определитель равен нулю.")
// 	}
// }
