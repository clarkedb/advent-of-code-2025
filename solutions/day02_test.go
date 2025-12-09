package solutions

import (
	"strings"
	"testing"
)

func TestDay02Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			want:  "1227775554",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day02Part1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day02Part1() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDay02Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			want:  "4174379265",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day02Part2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day02Part2() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDigitsFromInteger(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{12, []int{1, 2}},
		{123, []int{1, 2, 3}},
		{100, []int{1, 0, 0}},
		{987654321, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		got := digitsFromInteger(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("digitsFromInteger(%d) = %v, want %v", tt.input, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("digitsFromInteger(%d) = %v, want %v", tt.input, got, tt.want)
				break
			}
		}
	}
}

func TestContainsNRepeatedSubSequences_Int(t *testing.T) {
	type args[T comparable] struct {
		seq []T
		n   int
	}
	tests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "empty sequence",
			args: args[int]{seq: []int{}, n: 2},
			want: false,
		},
		{
			name: "n greater than length",
			args: args[int]{seq: []int{1, 2}, n: 3},
			want: false,
		},
		{
			name: "n is zero",
			args: args[int]{seq: []int{1, 2}, n: 0},
			want: false,
		},
		{
			name: "length not divisible by n",
			args: args[int]{seq: []int{1, 2, 3}, n: 2},
			want: false,
		},
		{
			name: "all elements same, n=2",
			args: args[int]{seq: []int{5, 5, 5, 5}, n: 2},
			want: true,
		},
		{
			name: "two repeated subsequences",
			args: args[int]{seq: []int{1, 2, 1, 2}, n: 2},
			want: true,
		},
		{
			name: "three repeated subsequences",
			args: args[int]{seq: []int{3, 4, 3, 4, 3, 4}, n: 3},
			want: true,
		},
		{
			name: "not repeated subsequences",
			args: args[int]{seq: []int{1, 2, 3, 4}, n: 2},
			want: false,
		},
		{
			name: "n=1 always true if n==len(seq)",
			args: args[int]{seq: []int{7}, n: 1},
			want: true,
		},
		{
			name: "n=1, longer seq",
			args: args[int]{seq: []int{1, 2, 3}, n: 1},
			want: true,
		},
		{
			name: "n equals length, all same",
			args: args[int]{seq: []int{9, 9, 9, 9}, n: 4},
			want: true,
		},
		{
			name: "n equals length, not all same",
			args: args[int]{seq: []int{1, 2, 3, 4}, n: 4},
			want: false,
		},
		{
			name: "sequence with partial repeats",
			args: args[int]{seq: []int{1, 2, 1, 3}, n: 2},
			want: false,
		},
		{
			name: "sequence with n=3, not repeated",
			args: args[int]{seq: []int{1, 2, 3, 1, 2, 4}, n: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNRepeatedSubSequences(tt.args.seq, tt.args.n); got != tt.want {
				t.Errorf("containsNRepeatedSubSequences(%v, %d) = %v, want %v", tt.args.seq, tt.args.n, got, tt.want)
			}
		})
	}
}

func TestContainsNRepeatedSubSequences_Strings(t *testing.T) {
	type args[T comparable] struct {
		seq []T
		n   int
	}
	tests := []struct {
		name string
		args args[string]
		want bool
	}{
		{
			name: "empty string sequence",
			args: args[string]{seq: []string{}, n: 2},
			want: false,
		},
		{
			name: "two repeated string subsequences",
			args: args[string]{seq: []string{"hello", "world", "hello", "world"}, n: 2},
			want: true,
		},
		{
			name: "three repeated string subsequences",
			args: args[string]{seq: []string{"a", "b", "a", "b", "a", "b"}, n: 3},
			want: true,
		},
		{
			name: "not repeated string subsequences",
			args: args[string]{seq: []string{"cat", "dog", "fish", "bird"}, n: 2},
			want: false,
		},
		{
			name: "single string repeated",
			args: args[string]{seq: []string{"test"}, n: 1},
			want: true,
		},
		{
			name: "all same strings",
			args: args[string]{seq: []string{"same", "same", "same", "same"}, n: 4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNRepeatedSubSequences(tt.args.seq, tt.args.n); got != tt.want {
				t.Errorf("containsNRepeatedSubSequences(%v, %d) = %v, want %v", tt.args.seq, tt.args.n, got, tt.want)
			}
		})
	}
}

func TestContainsNRepeatedSubSequences_Runes(t *testing.T) {
	type args[T comparable] struct {
		seq []T
		n   int
	}
	tests := []struct {
		name string
		args args[rune]
		want bool
	}{
		{
			name: "empty rune sequence",
			args: args[rune]{seq: []rune{}, n: 2},
			want: false,
		},
		{
			name: "two repeated rune subsequences",
			args: args[rune]{seq: []rune{'a', 'b', 'a', 'b'}, n: 2},
			want: true,
		},
		{
			name: "three repeated rune subsequences",
			args: args[rune]{seq: []rune{'x', 'y', 'z', 'x', 'y', 'z', 'x', 'y', 'z'}, n: 3},
			want: true,
		},
		{
			name: "not repeated rune subsequences",
			args: args[rune]{seq: []rune{'a', 'b', 'c', 'd'}, n: 2},
			want: false,
		},
		{
			name: "single rune",
			args: args[rune]{seq: []rune{'z'}, n: 1},
			want: true,
		},
		{
			name: "unicode runes repeated",
			args: args[rune]{seq: []rune{'😀', '😃', '😀', '😃'}, n: 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNRepeatedSubSequences(tt.args.seq, tt.args.n); got != tt.want {
				t.Errorf("containsNRepeatedSubSequences(%v, %d) = %v, want %v", tt.args.seq, tt.args.n, got, tt.want)
			}
		})
	}
}
