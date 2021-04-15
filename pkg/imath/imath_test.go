package imath

import (
	"testing"
)

func TestAbs(t *testing.T) {
	var tests = []struct{ value, result int }{
		{value: -10, result: 10},
		{value: 3, result: 3},
		{value: 0, result: 0},
	}

	for _, test := range tests {
		valueGot := Abs(test.value)
		if valueGot != test.result {
			t.Errorf("Got %d, expecting %d", valueGot, test.result)
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct{ a, b, result int }{
		{a: -10, b: 0, result: 0},
		{a: 0, b: 10, result: 10},
		{a: 2, b: 2, result: 2},
		{a: -100, b: -50, result: -50},
	}

	for _, test := range tests {
		valueGot := Max(test.a, test.b)
		if valueGot != test.result {
			t.Errorf("Got %d, expecting %d", valueGot, test.result)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct{ a, b, result int }{
		{a: -10, b: 0, result: -10},
		{a: 0, b: 10, result: 0},
		{a: 2, b: 2, result: 2},
		{a: -100, b: -50, result: -100},
	}

	for _, test := range tests {
		valueGot := Min(test.a, test.b)
		if valueGot != test.result {
			t.Errorf("Got %d, expecting %d", valueGot, test.result)
		}
	}
}
