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
			if isRepeatingSequenceInt(id) {
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
	var lines int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	// TODO: implement solution
	return fmt.Sprintf("processed %d lines", lines), nil
}

func isRepeatingSequenceInt(n int) bool {
	digits := make([]int, 0, 10)
	pos := n
	for pos > 0 {
		digits = append(digits, pos%10)
		pos /= 10
	}

	// reverse to get correct order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if len(digits)%2 != 0 || len(digits) < 2 {
		return false
	}

	mid := len(digits) / 2
	for i := range mid {
		if digits[i] != digits[mid+i] {
			return false
		}
	}
	return true
}
