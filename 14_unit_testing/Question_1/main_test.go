package main

import (
	"testing"
)

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		length, width int
		exp           string
	}{
		{4, 5, "even rectangle"},
		{7, 3, "odd rectangle"},
		{6, 8, "even rectangle"},
		{1, 1, "odd rectangle"},
	}

	for _, test := range tests {
		result := RectangleArea(test.length, test.width)
		if result != test.exp {
			t.Errorf("RectangleArea(%d, %d) = %s; expected %s", test.length, test.width, result, test.exp)
		}
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		length, width int
		exp           bool
	}{
		{4, 6, true},
		{7, 3, true},
		{10, 5, true},
		{1, 1, false},
	}

	for _, test := range tests {
		result := RectanglePerimeter(test.length, test.width)
		if result != test.exp {
			t.Errorf("RectanglePerimeter(%d, %d) = %v; expected %v", test.length, test.width, result, test.exp)
		}
	}
}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		side int
		exp  string
	}{
		{4, "even square"},
		{7, "odd square"},
		{10, "even square"},
		{1, "odd square"},
	}

	for _, test := range tests {
		result := SquareArea(test.side)
		if result != test.exp {
			t.Errorf("SquareArea(%d) = %s; expected %s", test.side, result, test.exp)
		}
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		side int
		exp  bool
	}{
		{4, false},
		{10, true},
		{9, false},
		{11, true},
	}

	for _, test := range tests {
		result := SquarePerimeter(test.side)
		if result != test.exp {
			t.Errorf("SquarePerimeter(%d) = %v; expected %v", test.side, result, test.exp)
		}
	}
}
