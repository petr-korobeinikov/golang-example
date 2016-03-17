package foundation

import (
	"github.com/pkorobeinikov/golang-example/util"
	"reflect"
	"testing"
)

func TestQueueLen(t *testing.T) {
	// Arrange
	cases := []struct {
		expected int
		given    Queue
	}{
		{0, Queue{}},
		{3, Queue{1, 2, 3}},
	}

	for _, c := range cases {
		// Act
		actual := c.given.Len()

		// Assert
		if actual != c.expected {
			t.Errorf("%v.Len() == %v, expected %v", c.given, actual, c.expected)
		}
	}
}

func TestQueueEnqueue(t *testing.T) {
	// Arrange
	cases := []struct {
		expected Queue
		given    Queue
		elements []interface{}
	}{
		{Queue{}, Queue{}, util.MakeInterfaceSliceFromIntSlice([]int{})},
		{Queue{1, 2, 3}, Queue{}, util.MakeInterfaceSliceFromIntSlice([]int{1, 2, 3})},
	}

	for _, c := range cases {
		// Act
		for _, elem := range c.elements {
			c.given.Enqueue(elem)
		}
		actual := c.given

		// Assert
		if !reflect.DeepEqual(c.expected, actual) {
			fmt := "%v == %v after enqueueing %v"
			t.Errorf(fmt, c.expected, actual, c.elements)
		}
	}
}

func TestQueueDequeue(t *testing.T) {
	// Arrange
	cases := []struct {
		expected Queue
		given    Queue
	}{
		{Queue{}, Queue{1}},
		{Queue{2, 3}, Queue{1, 2, 3}},
	}

	for _, c := range cases {
		// Act
		saveGiven := c.given
		dequeued, _ := c.given.Dequeue()

		// Assert
		if !reflect.DeepEqual(c.expected, c.given) {
			fmt := "%v.Dequeue() == %v, expected %v, dequeued %v"
			t.Errorf(fmt, saveGiven, c.given, c.expected, dequeued)
		}
	}
}

func TestQueueDequeueFromEmptyQueue(t *testing.T) {
	// Arrange
	sut := &Queue{}

	// Act
	_, err := sut.Dequeue()

	// Assert
	if err == nil {
		t.Errorf("Dequeuing from an empty queue must lead to an error.")
	}
}
