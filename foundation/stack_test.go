package foundation

import (
	"errors"
	"github.com/pkorobeinikov/golang-example/util"
	"reflect"
	"testing"
)

func TestStackLen(t *testing.T) {
	// Arrange
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
			t.Errorf("%v.Len() == %d, expected %d", c.given, actual, c.expected)
		}
	}
}

func TestStackPush(t *testing.T) {
	// Arrange
	cases := []struct {
		expected Stack
		given    Stack
		elements []interface{}
	}{
		{Stack{}, Stack{}, util.MakeInterfaceSliceFromIntSlice([]int{})},
		{Stack{1, 2, 3}, Stack{}, util.MakeInterfaceSliceFromIntSlice([]int{1, 2, 3})},
		{Stack{1, 2, 3, 4, 5}, Stack{}, util.MakeInterfaceSliceFromIntSlice([]int{1, 2, 3, 4, 5})},
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

func TestStackPop(t *testing.T) {
	// Arrange
	cases := []struct {
		expected interface{}
		err      error
		given    Stack
	}{
		{nil, errors.New("Unable to pop element from empty stack."), Stack{}},
		{1, nil, Stack{1}},
		{3, nil, Stack{1, 2, 3}},
	}

	for _, c := range cases {
		// Act
		saveGiven := c.given
		actual, e := c.given.Pop()

		// Assert
		errorsNotEqual := !(c.err != nil && e != nil && c.err != e)
		if c.expected != actual && errorsNotEqual {
			fmt := "%v.Pop() == %v, %v, expected %v, %v"
			t.Errorf(fmt, saveGiven, actual, e, c.expected, c.err)
		}
	}
}

func TestStackTop(t *testing.T) {
	// Arrange
	cases := []struct {
		expected interface{}
		err      error
		given    Stack
	}{
		{nil, errors.New("There is no top element in empty stack."), Stack{}},
		{1, nil, Stack{1}},
		{3, nil, Stack{1, 2, 3}},
	}

	for _, c := range cases {
		// Act
		actual, e := c.given.Top()

		// Assert
		errorsNotEqual := !(c.err != nil && e != nil && c.err != e)
		if c.expected != actual && errorsNotEqual {
			fmt := "%v.Top() == %v, %v, expected %v, %v"
			t.Errorf(fmt, c.given, actual, e, c.expected, c.err)
		}
	}
}
