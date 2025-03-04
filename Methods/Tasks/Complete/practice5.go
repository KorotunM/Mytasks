package main

import (
	"fmt"
	"math"
)

// Функция для нахождения нормы вектора
func norm(vec []float64) float64 {
	sum := 0.0
	for _, v := range vec {
		sum += v * v
	}
	return math.Sqrt(sum)
}

// Метод Якоби
func jacobiMethod(A [][]float64, b []float64, x0 []float64, tol float64, maxIter int) ([]float64, []float64) {
	n := len(A)
	xNew := make([]float64, n)
	rNorms := make([]float64, 0)

	for iter := 0; iter < maxIter; iter++ {
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j {
					sum += A[i][j] * x0[j]
				}
			}
			xNew[i] = (b[i] - sum) / A[i][i]
		}

		// Рассчитываем невязку
		r := make([]float64, n)
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				sum += A[i][j] * xNew[j]
			}
			r[i] = sum - b[i]
		}

		// Норма невязки
		normR := norm(r)
		rNorms = append(rNorms, normR)

		// Проверка сходимости
		if normR < tol {
			fmt.Printf("Якоби метод сошелся за %d итераций\n", iter+1)
			return xNew, rNorms
		}

		copy(x0, xNew)
	}

	fmt.Println("Якоби метод не сошелся")
	return xNew, rNorms
}

// Метод Зейделя
// Метод Зейделя
// func gaussSeidelMethod(A [][]float64, b []float64, x0 []float64, tol float64, maxIter int) ([]float64, []float64) {
// 	n := len(A)
// 	x := make([]float64, n)
// 	copy(x, x0)
// 	rNorms := make([]float64, 0)

// 	for iter := 0; iter < maxIter; iter++ {
// 		for i := 0; i < n; i++ {
// 			sum := 0.0
// 			for j := 0; j < n; j++ {
// 				if i != j {
// 					sum += A[i][j] * x[j]
// 				}
// 			}
// 			x[i] = (b[i] - sum) / A[i][i]
// 		}

// 		// Рассчитываем невязку
// 		r := make([]float64, n)
// 		for i := 0; i < n; i++ {
// 			sum := 0.0
// 			for j := 0; j < n; j++ {
// 				sum += A[i][j] * x[j]
// 			}
// 			r[i] = sum - b[i]
// 		}

// 		// Норма невязки
// 		normR := norm(r)
// 		rNorms = append(rNorms, normR) // Накапливаем нормы на каждой итерации

// 		// Проверка сходимости
// 		if normR < tol {
// 			fmt.Printf("Метод Зейделя сошелся за %d итераций\n", iter+1)
// 			// Не делаем return, а только выходим из цикла
// 			break
// 		}
// 	}

// 	// Возвращаем результат и накопленные нормы
// 	return x, rNorms
// }

func gaussSeidelMethod(a [][]float64, b []float64, x []float64, tolerance float64, maxIterations int) ([]float64, []float64) {
	n := len(a)
	norms := []float64{}

	for k := 0; k < maxIterations; k++ {
		norm := 0.0
		for i := 0; i < n; i++ {
			sum := b[i]
			for j := 0; j < n; j++ {
				if j != i {
					sum -= a[i][j] * x[j]
				}
			}
			newXi := sum / a[i][i]
			norm += math.Pow(newXi-x[i], 2)
			x[i] = newXi
		}

		// Вычисляем норму невязки
		norms = append(norms, math.Sqrt(norm))

		// Проверяем на сходимость
		if math.Sqrt(norm) < tolerance {
			fmt.Printf("Метод Зейделя сошелся за %d итераций\n", k+1)
			break
		}
	}

	return x, norms
}

func Task5() {
	// Пример СЛАУ
	A := [][]float64{
		{12.14, 1.32, -0.78, -2.75},
		{-0.89, 16.75, 1.88, -1.55},
		{2.65, -1.27, -15.64, -0.64},
		{2.44, 1.52, 1.93, -11.43},
	}
	b := []float64{14.78, -12.14, -11.65, 4.26}
	x0 := []float64{0, 0, 0, 0} // Начальное приближение
	tol := 1e-4                 // Заданная точность
	maxIter := 1000             // Максимальное количество итераций

	// Решение методом Якоби
	fmt.Println("Метод Якоби:")
	xJacobi, rNormsJacobi := jacobiMethod(A, b, x0, tol, maxIter)
	fmt.Printf("Решение: %v\n", xJacobi)

	A1 := [][]float64{
		{12.14, 1.32, -0.78, -2.75},
		{-0.89, 16.75, 1.88, -1.55},
		{2.65, -1.27, -15.64, -0.64},
		{2.44, 1.52, 1.93, -11.43},
	}
	b1 := []float64{14.78, -12.14, -11.65, 4.26}
	x01 := []float64{0, 0, 0, 0}

	// Решение методом Зейделя
	fmt.Println("Метод Зейделя:")
	xSeidel, rNormsSeidel := gaussSeidelMethod(A1, b1, x01, tol, maxIter)
	fmt.Printf("Решение: %v\n", xSeidel)

	//Вывод значений нормы невязки (для построения графика)
	fmt.Println("Норма невязки для метода Якоби:", rNormsJacobi)
	fmt.Println("Норма невязки для метода Зейделя:", rNormsSeidel)
}
