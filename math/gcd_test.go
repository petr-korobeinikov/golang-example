package math

import "testing"

// Arrange
var cases = []struct {
	expected, a, b int
}{
	{1, 4, 3},
	{10, 40, 30},
}

var invalidCases = []struct {
	expected, a, b int
}{
	{2, 4, 3},
	{5, 40, 30},
}

func actAndAssert(t *testing.T, fp func(int, int) int) {
	for _, c := range cases {
		// Act
		actual := fp(c.a, c.b)

		// Assert
		if actual != c.expected {
			t.Errorf("GCD(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func actAndAssertWithInvalidData(t *testing.T, fp func(int, int) int) {
	for _, c := range invalidCases {
		// Act
		actual := fp(c.a, c.b)

		// Assert
		if actual == c.expected {
			t.Errorf("GCD(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func TestGCDEuclidean(t *testing.T) {
	actAndAssert(t, GCDEuclidean)
}

func TestGCDReminderRecursive(t *testing.T) {
	actAndAssert(t, GCDReminderRecursive)
}

func TestGCDReminder(t *testing.T) {
	actAndAssert(t, GCDReminder)
}

func TestGCDEuclideanWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, GCDEuclidean)
}

func TestGCDReminderRecursiveWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, GCDReminderRecursive)
}

func TestGCDReminderWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, GCDReminder)
}
