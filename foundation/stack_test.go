package foundation

import (
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
			fmt := "%v == %v after pushing %v"
			t.Errorf(fmt, c.expected, actual, c.elements)
		}
	}
}

func TestStackPop(t *testing.T) {
	// Arrange
	cases := []struct {
		expected interface{}
		given    Stack
	}{
		{1, Stack{1}},
		{3, Stack{1, 2, 3}},
	}

	for _, c := range cases {
		// Act
		saveGiven := c.given
		actual, _ := c.given.Pop()

		// Assert
		if c.expected != actual {
			fmt := "%v.Pop() == %v, expected %v"
			t.Errorf(fmt, saveGiven, actual, c.expected)
		}
	}
}

func TestStackPopOnEmptyStack(t *testing.T) {
	// Arrange
	sut := &Stack{}

	// Act
	_, err := sut.Pop()

	// Assert
	if err == nil {
		t.Errorf("Pop on an empty stack must lead to an error")
	}
}

func TestStackTop(t *testing.T) {
	// Arrange
	cases := []struct {
		expected interface{}
		given    Stack
	}{
		{1, Stack{1}},
		{3, Stack{1, 2, 3}},
	}

	for _, c := range cases {
		// Act
		actual, _ := c.given.Top()

		// Assert
		if c.expected != actual {
			fmt := "%v.Top() == %v, expected %v"
			t.Errorf(fmt, c.given, actual, c.expected)
		}
	}
}

func TestStackTopOnEmptyStack(t *testing.T) {
	// Arrange
	sut := &Stack{}

	// Act
	_, err := sut.Top()

	// Assert
	if err == nil {
		t.Errorf("Top on an empty stack must lead to an error")
	}
}
