package main

import "testing"

// func TestAdd(t *testing.T) {
// 	got := Add(2, 3)
// 	expected := 5
//
// 	if got != expected {
// 		t.Errorf("Add(2, 3) = %d; expected %d", got, expected)
// 	}
// }

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
		{"mixed", -2, 5, 3},
		{"zero", 0, 0, 0},
		{"wrong", 1, 1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
