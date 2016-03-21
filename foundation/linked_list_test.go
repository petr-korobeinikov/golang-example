package foundation

import "testing"

func TestLinkedListLen(t *testing.T) {
	// Arrange
	cases := []struct {
		expected int
		given    LinkedListNode
	}{
		{1, LinkedListNode{}},
		{1, LinkedListNode{1, nil}},
		{3, LinkedListNode{1, &LinkedListNode{2, &LinkedListNode{3, nil}}}},
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
