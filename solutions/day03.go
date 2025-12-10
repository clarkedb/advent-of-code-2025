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
		max, err := maxTwoDigitSubstring(cell)
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

func maxTwoDigitSubstring(s string) (int, error) {
	n := len(s)
	var max int

	for i := range n {
		for j := i + 1; j < n; j++ {
			v, err := strconv.Atoi(fmt.Sprintf("%c%c", s[i], s[j]))
			if err != nil {
				return 0, err
			}

			if v > max {
				max = v
			}
		}
	}

	return max, nil
}
