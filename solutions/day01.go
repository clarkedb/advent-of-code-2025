package solutions

import (
	"fmt"
	"strings"
)

func init() {
	Register(1, 1, day01Part1)
	Register(1, 2, day01Part2)
}

func day01Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// TODO: implement solution
	return fmt.Sprintf("processed %d lines", len(lines)), nil
}

func day01Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// TODO: implement solution
	return fmt.Sprintf("processed %d lines", len(lines)), nil
}
