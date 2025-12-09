package solutions

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	iio "github.com/clarkedb/advent-of-code-2025/internal/io"
)

func init() {
	Register(2, 1, day02Part1)
	Register(2, 2, day02Part2)
}

func day02Part1(input io.Reader) (string, error) {
	var invalidIDSum int
	scanner := bufio.NewScanner(input)
	scanner.Split(iio.ScanCommas)
	for scanner.Scan() {
		idRange := strings.Split(scanner.Text(), "-")
		if len(idRange) != 2 {
			return "", fmt.Errorf("invalid range: %q", scanner.Text())
		}
		rangeStart, err := strconv.Atoi(strings.TrimSpace(idRange[0]))
		if err != nil {
			return "", fmt.Errorf("invalid range start: %q", idRange[0])
		}
		rangeEnd, err := strconv.Atoi(strings.TrimSpace(idRange[1]))
		if err != nil {
			return "", fmt.Errorf("invalid range end: %q", idRange[1])
		}

		for id := rangeStart; id <= rangeEnd; id++ {
			if isRepeatingSquenceTwice(id) {
				invalidIDSum += id
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", invalidIDSum), nil
}

func day02Part2(input io.Reader) (string, error) {
	var invalidIDSum int
	scanner := bufio.NewScanner(input)
	scanner.Split(iio.ScanCommas)
	for scanner.Scan() {
		idRange := strings.Split(scanner.Text(), "-")
		if len(idRange) != 2 {
			return "", fmt.Errorf("invalid range: %q", scanner.Text())
		}
		rangeStart, err := strconv.Atoi(strings.TrimSpace(idRange[0]))
		if err != nil {
			return "", fmt.Errorf("invalid range start: %q", idRange[0])
		}
		rangeEnd, err := strconv.Atoi(strings.TrimSpace(idRange[1]))
		if err != nil {
			return "", fmt.Errorf("invalid range end: %q", idRange[1])
		}

		for id := rangeStart; id <= rangeEnd; id++ {
			if isRepeatingSequence(id) {
				invalidIDSum += id
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", invalidIDSum), nil
}

func digitsFromInteger(z int) []int {
	digits := make([]int, 0, 10)
	pos := z
	for pos > 0 {
		digits = append(digits, pos%10)
		pos /= 10
	}

	// reverse to get correct order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

func isRepeatingSquenceTwice(z int) bool {
	if z < 11 {
		return false
	}

	digits := digitsFromInteger(z)
	return containsNRepeatedSubSequences(digits, 2)
}

func isRepeatingSequence(z int) bool {
	if z < 11 {
		return false
	}

	digits := digitsFromInteger(z)

	for n := 2; n <= len(digits); n++ {
		if containsNRepeatedSubSequences(digits, n) {
			return true
		}
	}
	return false
}

func containsNRepeatedSubSequences[T comparable](seq []T, n int) bool {
	m := len(seq)
	if n > m || n < 1 || m%n != 0 {
		return false
	}

	subSeqLen := m / n
	for i := range subSeqLen {
		v := seq[i]
		for j := 1; j < n; j++ {
			if seq[i+j*subSeqLen] != v {
				return false
			}
		}
	}

	return true
}
