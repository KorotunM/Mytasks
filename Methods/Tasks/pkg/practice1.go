package pkg

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func fac(n float64) float64 {
	var (
		s float64 = 1
		i float64 = 1
	)
	for ; i <= n; i++ {
		s *= i
	}
	return s
}

func Erf(x float64) float64 {
	var sum float64 = 0
	var temp float64

	for n := 0; n < 100000; n++ {
		temp = (math.Pow(-1, float64(n)) * math.Pow(x, float64(2*n+1))) / (float64(fac(float64(n))) * float64(2*n+1))
		sum += temp

		if math.Abs(temp) < 1e-16 {
			break
		}
	}

	ans := sum * 2 / math.Sqrt(math.Pi)
	return ans
}

func Task2() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var (
		sum  float64
		x    float64 = 0
		k    float64
		temp float64
	)

	fmt.Fprintln(out, "Значения для первой суммы: ")
	for ; x <= 1; x += 0.1 {
		sum = 0
		for k = 1; ; k++ {
			temp = 1 / (k * (k + x))
			sum += temp

			if temp < 5e-9 {
				break
			}
		}
		fmt.Fprintf(out, "для x = %f: %f\n", x, sum)
	}

	fmt.Fprintln(out, "Значения для второй суммы: ")
	for x = 0; x <= 1; x += 0.1 {
		sum = 0
		for k = 1; ; k++ {
			temp = 1/k - 1/(k+x)
			sum += temp

			if temp < 5e-9 {
				break
			}
		}
		fmt.Fprintf(out, "для x = %f: %f\n", x, sum)
	}

	fmt.Fprintln(out, "Значения для третьей суммы: ")
	for x = 0; x <= 1; x += 0.1 {
		sum = 0
		for k = 1; ; k++ {
			temp = 1/(k*(k+x)) - 1/(k*(k+1))
			sum += temp

			if temp < 5e-9 {
				break
			}
		}
		fmt.Fprintf(out, "для x = %f: %f\n", x, sum)
	}
}

func Task3() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		sum  float64
		x    float64
		k    float64 = 1
		temp float64
		cnt  int
	)
	fmt.Fscan(in, &x)
	fmt.Fprintln(out, "Общее решение: ")
	for ; ; k++ {
		temp = 1/math.Sqrt(k*k*k+x) - 1/math.Sqrt(k*k*k-x)
		sum += temp

		if math.Abs(temp) < 3e-8 {
			cnt = int(k - 1)
			break
		}
	}
	fmt.Fprintf(out, "для x = %f: %f\n", x, sum)

	fmt.Fprintln(out, "--Пункт А: ")
	var temp_first float64
	for x = -0.9; x < 0.9; x += 0.1 {
		sum = 0
		for k = 1; ; k++ {
			temp_first = 1/math.Sqrt(k*k*k+x) - 1/math.Sqrt(k*k*k-x)
			sum += temp_first

			if math.Abs(temp_first) < 3e-8 {
				break
			}
		}
		fmt.Fprintf(out, "для x = %f: %f\n", x, sum)
	}
	fmt.Fprintln(out, "Мы видим что значения стремятся к нулю => ряд сходится")

	fmt.Fprintf(out, "--Пункт Б:\nКоличество = %d\n", cnt)
	fmt.Fprintf(out, "Пункт В:\nПримерное время ~ %d микросекунд\n", cnt*500)
	fmt.Fprintln(out, "Пункт Д:\nДля значения 0.9999999 по стандарту IEEE 754 не хватает места, \nпоэтому число округляется до ближайшего целого т.е. 1")

}

func Task4() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var (
		sum  float64 = 0
		temp float64
		n    float64 = 1
		cnt  int     = 0
	)
	for ; ; n++ {
		temp = 1 / (n*n + 1)
		sum += temp

		if temp < 1e-10 {
			cnt = int(n - 1)
			break
		}
	}
	fmt.Fprintf(out, "Ряд1  = %f\nКоличество1: %d\n", sum, cnt)
	sum = 0
	n = 1
	for ; n < 10000; n++ {
		temp = 1 / (math.Pow(n, 4) * (n*n + 1))
		sum += temp
		if temp < 1e-10 {
			cnt = int(n - 1)
			break
		}
	}
	sum += math.Pi*math.Pi/6 - math.Pow(3.14519, float64(4))/90
	fmt.Fprintf(out, "Ряд2  = %f\nКоличество2: %d\n", sum, cnt)
}
