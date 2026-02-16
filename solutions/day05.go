package solutions

import (
	"bufio"
	"fmt"
	"io"
	"slices"
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

	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				count++
				break
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func day05Part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	ranges := make([][2]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		var a, b int
		_, err := fmt.Sscanf(line, "%d-%d", &a, &b)
		if err != nil {
			return "", fmt.Errorf("error reading range: %w", err)
		}
		ranges = append(ranges, [2]int{a, b})
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	slices.SortFunc(ranges, func(a, b [2]int) int { return a[0] - b[0] })

	merged := make([][2]int, 0, len(ranges))
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		m := merged[len(merged)-1]
		if r[0] > m[1] {
			merged = append(merged, r)
		} else if r[1] > m[1] {
			merged[len(merged)-1][1] = r[1]
		}
	}

	count := 0
	for _, r := range merged {
		count += r[1] - r[0] + 1
	}

	return fmt.Sprintf("%d", count), nil
}
