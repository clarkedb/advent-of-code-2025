package solutions

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func init() {
	Register(3, 1, day03Part1)
	Register(3, 2, day03Part2)
}

func day03Part1(input io.Reader) (string, error) {
	var sum int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cell := scanner.Text()
		max, err := maxNDigitSubstring(cell, 2)
		if err != nil {
			return "", err
		}
		sum += max
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", sum), nil
}

func day03Part2(input io.Reader) (string, error) {
	var sum int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cell := scanner.Text()
		max, err := maxNDigitSubstring(cell, 12)
		if err != nil {
			return "", err
		}
		sum += max
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", sum), nil
}

func maxNDigitSubstring(s string, n int) (int, error) {
	if len(s) < n {
		return 0, nil
	}

	// greedy approach: build the largest n-digit number
	result := make([]byte, n)
	digitIndex := 0

	for pos := range n {
		// find the largest digit we can use at this position
		// while ensuring we have enough digits left for remaining positions
		maxDigit := byte('0')
		maxDigitIndex := -1

		// Look ahead: we need (n - pos - 1) more digits after this one
		remainingPositions := n - pos - 1

		for i := digitIndex; i <= len(s)-1-remainingPositions; i++ {
			if s[i] > maxDigit {
				maxDigit = s[i]
				maxDigitIndex = i
			}
		}

		if maxDigitIndex == -1 {
			// This shouldn't happen if our logic is correct
			return 0, fmt.Errorf("failed to find digit at position %d", pos)
		}

		result[pos] = maxDigit
		digitIndex = maxDigitIndex + 1
	}

	val, err := strconv.Atoi(string(result))
	if err != nil {
		return 0, err
	}

	return val, nil
}
