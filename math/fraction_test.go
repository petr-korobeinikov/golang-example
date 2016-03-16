package math

import "testing"

func TestString(t *testing.T) {
	// Arrange
	cases := []struct {
		expected string
		given    Fraction
	}{
		{"1/1", Fraction{1, 1}},
		{"1/2", Fraction{1, 2}},
		{"1/5", Fraction{1, 5}},
	}

	for _, c := range cases {
		// Act
		actual := c.given.String()

		// Assert
		if actual != c.expected {
			fmt := "String representation of Fraction{%d, %d} == %v, expected %s"
			t.Errorf(fmt, c.given.Numerator, c.given.Denominator, actual, c.expected)
		}
	}
}

func TestStringWithInvalidData(t *testing.T) {
	// Arrange
	invalidCases := []struct {
		expected string
		given    Fraction
	}{
		{"1/5", Fraction{1, 1}},
	}

	for _, c := range invalidCases {
		// Act
		actual := c.given.String()

		// Assert
		if actual == c.expected {
			fmt := "String representation of Fraction{%d, %d} == %v, expected %s"
			t.Errorf(fmt, c.given.Numerator, c.given.Denominator, actual, c.expected)
		}
	}
}

func TestReduce(t *testing.T) {
	// Arrange
	cases := []struct {
		expected, given Fraction
	}{
		{Fraction{1, 1}, Fraction{10, 10}},
		{Fraction{1, 2}, Fraction{2, 4}},
		{Fraction{1, 5}, Fraction{2, 10}},
	}

	for _, c := range cases {
		// Act
		actual := c.given.Reduce()

		// Assert
		if actual != c.expected {
			t.Errorf("Fraction.Reduce(%v) = %v, expected %v", c.given, actual, c.expected)
		}
	}
}

func TestReduceWithInvalidData(t *testing.T) {
	// Arrange
	invalidCases := []struct {
		expected, given Fraction
	}{
		{Fraction{1, 10}, Fraction{10, 10}},
		{Fraction{1, 20}, Fraction{2, 4}},
		{Fraction{1, 50}, Fraction{2, 10}},
	}

	for _, c := range invalidCases {
		// Act
		actual := c.given.Reduce()

		// Assert
		if actual == c.expected {
			t.Errorf("Fraction.Reduce(%v) = %v, expected %v", c.given, actual, c.expected)
		}
	}
}

func TestMultiplyByNumber(t *testing.T) {
	// Arrange
	cases := []struct {
		expected, given Fraction
		m               int
	}{
		{Fraction{6, 7}, Fraction{3, 7}, 2},
		{Fraction{4, 2}, Fraction{1, 2}, 4},
	}

	for _, c := range cases {
		// Act
		actual := c.given.MultiplyByNumber(c.m)

		// Assert
		if actual != c.expected {
			t.Errorf("Fraction.MultiplyByNumber(%v) = %v, expected %v", c.given, actual, c.expected)
		}
	}
}

func TestMultiplyByNumberWithInvalidData(t *testing.T) {
	// Arrange
	invalidCases := []struct {
		expected, given Fraction
		m               int
	}{
		{Fraction{6, 7}, Fraction{3, 7}, 3},
		{Fraction{2, 2}, Fraction{1, 2}, 5},
	}

	for _, c := range invalidCases {
		// Act
		actual := c.given.MultiplyByNumber(c.m)

		// Assert
		if actual == c.expected {
			t.Errorf("Fraction.MultiplyByNumber(%v) = %v, expected %v", c.given, actual, c.expected)
		}
	}
}

func TestMultiplyByFraction(t *testing.T) {
	// Arrange
	cases := []struct {
		expected, a, b Fraction
	}{
		{Fraction{6, 35}, Fraction{3, 7}, Fraction{2, 5}},
		{Fraction{30, 36}, Fraction{3, 4}, Fraction{10, 9}},
	}

	for _, c := range cases {
		// Act
		actual := c.a.MultiplyByFraction(c.b)

		// Assert
		if actual != c.expected {
			t.Errorf("Fraction.MultiplyByFraction(%v) = %v, expected %v", c.a, actual, c.expected)
		}
	}
}

func TestMultiplyByFractionWithInvalidData(t *testing.T) {
	// Arrange
	invalidCases := []struct {
		expected, a, b Fraction
	}{
		{Fraction{16, 35}, Fraction{3, 7}, Fraction{2, 5}},
		{Fraction{130, 36}, Fraction{3, 4}, Fraction{10, 9}},
	}

	for _, c := range invalidCases {
		// Act
		actual := c.a.MultiplyByFraction(c.b)

		// Assert
		if actual == c.expected {
			t.Errorf("Fraction.MultiplyByFraction(%v) = %v, expected %v", c.a, actual, c.expected)
		}
	}
}
