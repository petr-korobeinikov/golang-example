package parenthesis

import "testing"

func TestIsMatchingPositive(t *testing.T) {
	cases := []string{
		"",

		"()",
		"{}",
		"[]",

		"[] [] []",
		"[] ({}) {}",

		"({})",
		"([])",
		"{()}",

		"((()))",
		"((({{{}}})))",
		"((({{{[[[]]]}}})))",

		"some data ( should be valid ( even ({{{ with [[[ different types ]]]} of }}) brackets))",
	}

	for _, c := range cases {
		if !IsMatching(c) {
			t.Errorf("IsMatching(%v) expected to be true", c)
		}
	}
}

func TestIsMatchingNegative(t *testing.T) {
	cases := []string{
		"(",
		"[",
		"{",

		")",
		"]",
		"}",

		"{(})",
		"[(}]",
		"([)]",

		"([",
		"[(",
		"{(",
	}

	for _, c := range cases {
		if IsMatching(c) {
			t.Errorf("IsMatching(%v) expected to be false", c)
		}
	}
}
