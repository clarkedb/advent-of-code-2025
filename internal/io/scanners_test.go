package io

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

// signature verification for bufio.SplitFunc
var _ bufio.SplitFunc = ScanCommas

func TestScanCommas(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "Standard list",
			input: "a,b,c,d",
			want:  []string{"a", "b", "c", "d"},
		},
		{
			name:  "Empty items",
			input: "apple,,banana",
			want:  []string{"apple", "", "banana"},
		},
		{
			name:  "No commas",
			input: "singleword",
			want:  []string{"singleword"},
		},
		{
			name:  "Empty input",
			input: "",
			want:  nil, // Scanner loop won't run, result slice stays nil/empty
		},
		{
			name:  "Trailing comma",
			input: "one,two,",
			// Note: The logic in ScanCommas behaves like ScanLines.
			// "one,two" returns "one", "two".
			// "one,two," effectively means the last item is empty,
			// but because of how EOF is handled, it might drop the final empty token
			// depending on strict implementation.
			// With the current implementation:
			// 1. "one"
			// 2. "two"
			// 3. Remainder is empty string at EOF -> returns 0, nil, nil.
			want: []string{"one", "two"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			scanner.Split(ScanCommas)

			var got []string
			for scanner.Scan() {
				got = append(got, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				t.Fatalf("scanner error: %v", err)
			}

			// handle nil vs empty slice comparison
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
