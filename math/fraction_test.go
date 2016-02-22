package math

import "testing"

func TestString(t *testing.T) {
	// Arrange
	cases := []struct {
		expected string
		given Fraction
	}{
		{"1/1", Fraction{1, 1}},
		{"1/2", Fraction{1, 2}},
		{"1/5", Fraction{1, 5}},
	}

	for _, c := range cases {
		// Act
		actual := c.given.String();

		// Assert
		if actual != c.expected {
			fmt := "String representation of Fraction{%d, %d} == %q, expected %s"
			t.Errorf(fmt, c.given.Numerator, c.given.Denominator, actual, c.expected)
		}
	}
}

func TestStringWithInvalidData(t *testing.T) {
	// Arrange
	invalidCases := []struct {
		expected string
		given Fraction
	}{
		{"1/5", Fraction{1, 1}},
	}

	for _, c := range invalidCases {
		// Act
		actual := c.given.String();

		// Assert
		if actual == c.expected {
			fmt := "String representation of Fraction{%d, %d} == %q, expected %s"
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
		actual := c.given.Reduce();

		// Assert
		if actual != c.expected {
			t.Errorf("Fraction.Reduce(%q) = %q, expected %q", c.given, actual, c.expected)
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
		actual := c.given.Reduce();

		// Assert
		if actual == c.expected {
			t.Errorf("Fraction.Reduce(%q) = %q, expected %q", c.given, actual, c.expected)
		}
	}
}
