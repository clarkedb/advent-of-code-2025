package solutions

import (
	"strings"
	"testing"
)

func TestDay03Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "987654321111111\n811111111111119\n234234234234278\n818181911112111",
			want:  "357",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day03Part1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day03Part1() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDay03Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "987654321111111\n811111111111119\n234234234234278\n818181911112111",
			want:  "3121910778619",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day03Part2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day03Part2() = %q, want %q", got, tt.want)
			}
		})
	}
}
