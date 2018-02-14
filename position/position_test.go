package position

import "testing"

func TestPosition(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		x     int
		want  int
	}{
		{"x is in array", []int{1, 2, 3, 4}, 2, 1},
		{"x is not in array", []int{1, 2, 3, 4}, 7, -1},
		{"x is the first element", []int{1, 2, 3, 4}, 1, 0},
		{"x is the last element", []int{1, 2, 3, 4}, 4, 3},
		{"nil array", nil, 1, -1},
		{"empty array", []int{}, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Position(tt.array, tt.x); got != tt.want {
				t.Errorf("Name: %s, Position() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
