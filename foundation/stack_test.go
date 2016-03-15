package foundation

import (
	"reflect"
	"testing"
)

func makeInterfaceSlice(in []int) (out []interface{}) {
	out = make([]interface{}, len(in))
	for i, d := range in {
		out[i] = d
	}
	return
}

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
		actual := c.given.Len()

		// Assert
		if actual != c.expected {
			t.Errorf("%q.Len() == %d, expected %d", c.given, actual, c.expected)
		}
	}
}

func TestPush(t *testing.T) {
	cases := []struct {
		expected Stack
		given    Stack
		elements []interface{}
	}{
		{Stack{}, Stack{}, makeInterfaceSlice([]int{})},
		{Stack{1, 2, 3}, Stack{}, makeInterfaceSlice([]int{1, 2, 3})},
		{Stack{1, 2, 3, 4, 5}, Stack{}, makeInterfaceSlice([]int{1, 2, 3, 4, 5})},
	}

	for _, c := range cases {
		// Act
		for _, elem := range c.elements {
			c.given.Push(elem)
		}
		actual := c.given

		// Assert
		if !reflect.DeepEqual(c.expected, actual) {
			fmt := "%v != %v after pushing %v"
			t.Errorf(fmt, c.expected, actual, c.elements)
		}
	}
}
