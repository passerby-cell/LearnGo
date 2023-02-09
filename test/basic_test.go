package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{12, 35, 39},
	}
	for _, test := range tests {
		if real := triangle(test.a, test.b); real != test.c {
			t.Errorf("triangle(%d, %d) = %d, want %d", test.a, test.b, real, test.c)

		}
	}
}
