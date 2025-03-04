package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Меняем местами строки в матрице
func swapRows(A *mat.Dense, i, j int) {
	r, c := A.Dims()
	if i < 0 || j < 0 || i >= r || j >= r {
		fmt.Println("Ошибка: индексы строк выходят за пределы.")
		return
	}

	// Меняем местами строки i и j
	for k := 0; k < c; k++ {
		temp := A.At(i, k)
		A.Set(i, k, A.At(j, k))
		A.Set(j, k, temp)
	}
}

// Приведение матрицы к ступенчатому виду (метод Гаусса)
func Gaus(A *mat.Dense, b *mat.VecDense) (*mat.Dense, *mat.VecDense) {
	r, c := A.Dims()
	temp := mat.NewDense(r, c+1, nil)

	// Создаем расширенную матрицу [A|b]
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			temp.Set(i, j, A.At(i, j))
		}
		temp.Set(i, c, b.AtVec(i))
	}

	// Прямой ход метода Гаусса (приведение к ступенчатому виду)
	for i := 0; i < r; i++ {
		// Поиск ведущего элемента
		if temp.At(i, i) == 0 {
			for k := i + 1; k < r; k++ {
				if temp.At(k, i) != 0 {
					swapRows(temp, i, k) // Меняем строки местами
					break
				}
			}
		}

		// Нормализация строки
		Max_fromColumn := temp.At(i, i)
		if Max_fromColumn != 0 {
			for j := 0; j <= c; j++ {
				temp.Set(i, j, temp.At(i, j)/Max_fromColumn)
			}
		}

		// Обнуление элементов ниже ведущего
		for k := i + 1; k < r; k++ {
			factor := temp.At(k, i)
			for j := 0; j <= c; j++ {
				temp.Set(k, j, temp.At(k, j)-factor*temp.At(i, j))
			}
		}
	}

	// Разделяем матрицу на A и b после приведения к ступенчатому виду
	newA := mat.NewDense(r, c, nil)
	newB := mat.NewVecDense(r, nil)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			newA.Set(i, j, temp.At(i, j))
		}
		newB.SetVec(i, temp.At(i, c))
	}

	return newA, newB
}

// Вывод общего решения системы
func findGeneralSolution(A *mat.Dense, b *mat.VecDense) {
	r, c := A.Dims()
	//fmt.Println("Матрица после приведения к ступенчатому виду:")
	//matPrint(A)
	//fmt.Println("Вектор правых частей:")
	//matPrintVec(b)

	// Выявление свободных переменных
	fmt.Println("Решение через свободные переменные:")
	for i := 0; i < r; i++ {
		foundLeading := false
		for j := 0; j < c; j++ {
			if A.At(i, j) != 0 {
				foundLeading = true
				fmt.Printf("x%d = ", j+1)
				for k := j + 1; k < c; k++ {
					if A.At(i, k) != 0 {
						fmt.Printf("- %.2ft%d ", A.At(i, k), k+1)
					}
				}
				fmt.Printf("+ %.2f\n", b.AtVec(i))
				break
			}
		}
		if !foundLeading {
			fmt.Printf("Свободная переменная: t%d\n", i+1)
		}
	}
}

func Task4_4_2() {
	A := mat.NewDense(3, 3, []float64{
		0.1, 0.2, 0.3,
		0.4, 0.5, 0.6,
		0.7, 0.8, 0.9,
	})

	b := mat.NewVecDense(3, []float64{
		0.1, 0.3, 0.5,
	})

	newA, newB := Gaus(A, b)

	newA.Set(2, 2, 0)
	newB.SetVec(2, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(" ", newA.At(i, j))
		}
		fmt.Printf("   %v\n", newB.AtVec(i))
	}

	findGeneralSolution(newA, newB)
}
