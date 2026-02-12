package solutions

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func init() {
	Register(5, 1, day05Part1)
	Register(5, 2, day05Part2)
}

func day05Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	ranges := make([][2]int, 0)
	ids := make([]int, 0)
	readRanges := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			readRanges = false
			continue
		}

		if readRanges {
			var a, b int
			_, err := fmt.Sscanf(line, "%d-%d", &a, &b)
			if err != nil {
				return "", fmt.Errorf("error reading range: %w", err)
			}
			ranges = append(ranges, [2]int{a, b})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return "", fmt.Errorf("error reading id: %w", err)
			}
			ids = append(ids, id)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// TODO: implement solution
	return fmt.Sprintf("processed %d ranges and %d ids", len(ranges), len(ids)), nil
}

func day05Part2(input io.Reader) (string, error) {
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
