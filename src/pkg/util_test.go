package pkg

import (
	"testing"
)

func TestTo(t *testing.T) {
	tests := []struct {
		condition bool
		t1        interface{}
		t2        interface{}
		expected  interface{}
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
		{true, "a", "b", "a"},
		{false, "a", "b", "b"},
	}

	for _, test := range tests {
		result := To(test.condition, test.t1, test.t2)
		if result != test.expected {
			t.Errorf("To(%v, %v, %v) = %v; want %v", test.condition, test.t1, test.t2, result, test.expected)
		}
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		condition bool
		t1        func() interface{}
		t2        func() interface{}
		expected  interface{}
	}{
		{true, func() interface{} { return 1 }, func() interface{} { return 2 }, 1},
		{false, func() interface{} { return 1 }, func() interface{} { return 2 }, 2},
		{true, func() interface{} { return "a" }, func() interface{} { return "b" }, "a"},
		{false, func() interface{} { return "a" }, func() interface{} { return "b" }, "b"},
	}

	for _, test := range tests {
		result := Do(test.condition, test.t1, test.t2)
		if result != test.expected {
			t.Errorf("Do(%v, t1, t2) = %v; want %v", test.condition, result, test.expected)
		}
	}
}
