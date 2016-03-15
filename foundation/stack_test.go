package foundation

import "testing"

func TestLen(t *testing.T) {
	cases := []struct {
		expected int
		given    Stack
	}{
		{0, Stack{}},
		{3, Stack{1, 2, 3}},
	}

	for _, c := range cases {
		// Act
		actual := c.given.Len();

		// Assert
		if actual != c.expected {
			t.Errorf("%q.Len() == %d, expected %d", c.given, actual, c.expected)
		}
	}
}
