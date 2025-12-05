package solutions

import (
	"strings"
	"testing"
)

func TestDay01Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			want:  "3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day01Part1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day01Part1() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDay01Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			want:  "6",
		},
		{
			name:  "multiple crossing example",
			input: "R1000",
			want:  "10",
		},
		{
			name:  "exact left crossing edge case",
			input: "L150",
			want:  "2",
		},
		{
			name:  "exact left zero case",
			input: "L50",
			want:  "1",
		},
		{
			name:  "toggle zero crossing",
			input: "L50\nR50\nL50\nR50\nL50",
			want:  "3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day01Part2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day01Part2() = %q, want %q", got, tt.want)
			}
		})
	}
}
