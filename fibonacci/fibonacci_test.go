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

func TestRecursively(t *testing.T) {
	for _, c := range cases {
		// Act
		actual := Recursively(c.n)

		// Assert
		if actual != c.fibn {
			t.Errorf("Recursively(%d) == %d, expected %d", c.n, actual, c.fibn)
		}
	}
}

func TestIteratively(t *testing.T) {
	for _, c := range cases {
		// Act
		actual := Iteratively(c.n)

		// Assert
		if actual != c.fibn {
			t.Errorf("Iteratively(%d) == %d, expected %d", c.n, actual, c.fibn)
		}
	}
}

func TestByGoldenRatio(t *testing.T) {
	for _, c := range cases {
		// Act
		actual := ByGoldenRatio(c.n)
		delta := 3

		// Assert
		eqWithDelta := (c.fibn - delta <= actual) && (actual <= c.fibn + delta)
		if !eqWithDelta {
			t.Errorf("ByGoldenRatio(%d) == %d±%d, expected %d", c.n, actual, delta, c.fibn)
		}
	}
}
