package fibonacci

import "math"

//    n =  0  1  2  3  4  5  6   7   8   9  10  11   12   13   14 ...
// fibn =  0  1  1  2  3  5  8  13  21  34  55  89  144  233  377 ...

// Recursively calculates Fibonacci numbers recursively.
func Recursively(n int) int {
	if n == 0 {
		return 0
	}

	if n < 3 {
		return 1
	}

	return Recursively(n - 1) + Recursively(n - 2)
}

// Iteratively calculates Fibonacci numbers iteratively.
func Iteratively(n int) int {
	fibs := []int{0, 1}

	for i := 2; i <= n; i++ {
		next := fibs[i - 1] + fibs[i - 2]
		fibs = append(fibs, next)
	}

	return fibs[n]
}

// ByGoldenRatio calculates Fibonacci numbers using golden ratio.
func ByGoldenRatio(n int) int {
	numerator := math.Pow(1.618034, float64(n)) - math.Pow(-0.618034, float64(n))
	denominator := math.Sqrt(5)

	return int(numerator / denominator);
}
