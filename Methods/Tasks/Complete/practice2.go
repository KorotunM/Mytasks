package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Task2_2() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c, x1, x2, d float64
	fmt.Fscan(in, &a, &b, &c)

	if a == 0 && b == 0 && c == 0 {
		fmt.Fprint(out, "Все комплексные числа удовлетворяют квадратному уравнению")
	} else if a == 0 && b == 0 && c != 0 {
		fmt.Fprint(out, "Ни одно комплексное число не удовлетворяет этому уравнению")
	} else if a == 0 && b != 0 {
		x1 = -c / b
		fmt.Fprintf(out, "Решение линейного уравнения: x = %.11f\n", x1)
		return
	}

	d = b*b - 4*a*c

	if d < 0 {
		sqrtKL := math.Sqrt(math.Abs(d))
		x_v1 := -b / (2 * a)
		x_k1 := sqrtKL / (2 * a)
		x_k2 := -sqrtKL / (2 * a)
		fmt.Fprintf(out, "Корни для комплексных:\nx1 = %.11f + (%.11f)i\nx2 = %.11f + (%.11f)i\n", x_v1, x_k1, x_v1, x_k2)
		return
	}

	sqrtD := math.Sqrt(d)

	x1 = (-b + sqrtD) / (2 * a)
	x2 = c / (a * x1)

	fmt.Fprintf(out, "Значения корней:\nx1 = %.11f\nx2 = %.11f", x1, x2)
}
