package util

import (
	"reflect"
	"testing"
)

func TestMakeInterfaceSlice(t *testing.T) {
	cases := []struct {
		given []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2, 3}},
	}

	for _, c := range cases {
		actual := MakeInterfaceSliceFromIntSlice(c.given)
		if "[]interface {}" != reflect.TypeOf(actual).String() {
			fmt := "MakeInterfaceSliceFromIntSlice(%v) is expected to return an []interface {}"
			t.Errorf(fmt, c.given, actual)
		}
	}
}
