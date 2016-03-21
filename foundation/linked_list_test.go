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
			t.Errorf("%v.Len() == %d, expected %d", c.given, actual, c.expected)
		}
	}
}

func TestLinkedListLast(t *testing.T) {
	// Arrange
	cases := []struct {
		expected interface{}
		given    LinkedListNode
	}{
		{nil, LinkedListNode{}},
		{1, LinkedListNode{1, nil}},
		{3, LinkedListNode{1, &LinkedListNode{2, &LinkedListNode{3, nil}}}},
	}

	for _, c := range cases {
		// Act
		actual := c.given.Last().Element

		// Assert
		if actual != c.expected {
			t.Errorf("%v.Last() == %d, expected %d", c.given, actual, c.expected)
		}
	}
}
