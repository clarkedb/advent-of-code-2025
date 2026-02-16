package solutions

import (
	"strings"
	"testing"
)

func TestDay05Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32",
			want:  "3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day05Part1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day05Part1() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDay05Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32",
			want:  "14",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day05Part2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day05Part2() = %q, want %q", got, tt.want)
			}
		})
	}
}
