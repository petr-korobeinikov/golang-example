package foundation

import "testing"

func TestLinkedListLen(t *testing.T) {
	// Arrange
	cases := []struct {
		expected int
		given    LinkedList
	}{
		{1, LinkedList{}},
		{1, LinkedList{1, nil}},
		{3, LinkedList{1, &LinkedList{2, &LinkedList{3, nil}}}},
	}

	for _, c := range cases {
		// Act
		actual := c.given.Len()

		// Assert
		if actual != c.expected {
			t.Errorf("%v.Len() == %d, expeced %d", c.given, actual, c.expected)
		}
	}
}
