package solutions

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	Register(1, 1, day01Part1)
	Register(1, 2, day01Part2)
}

func day01Part1(input io.Reader) (string, error) {
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

func day01Part2(input io.Reader) (string, error) {
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
