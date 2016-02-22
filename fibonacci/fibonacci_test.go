package fibonacci

import "testing"

// Arrange
var cases = []struct {
	fibn, n int
}{
	{0, 0},
	{1, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{5, 5},
	{21, 8},
	{377, 14},
	{14930352, 36},
}

func TestFibonacciRecursive(t *testing.T) {
	for _, c := range cases {
		// Act
		actual := FibonacciRecursive(c.n)

		// Assert
		if actual != c.fibn {
			t.Errorf("FibonacciRecursive(%d) == %d, expected %d", c.n, actual, c.fibn)
		}
	}
}

func TestFibonacciIterative(t *testing.T) {
	for _, c := range cases {
		// Act
		actual := FibonacciIterative(c.n)

		// Assert
		if actual != c.fibn {
			t.Errorf("FibonacciIterative(%d) == %d, expected %d", c.n, actual, c.fibn)
		}
	}
}

func TestFibonacciGoldenRatio(t *testing.T) {
	for _, c := range cases {
		// Act
		actual := FibonacciGoldenRatio(c.n)
		delta := 3

		// Assert
		eqWithDelta := (c.fibn - delta <= actual) && (actual <= c.fibn + delta)
		if !eqWithDelta {
			t.Errorf("FibonacciGoldenRatio(%d) == %d±%d, expected %d", c.n, actual, delta, c.fibn)
		}
	}
}
