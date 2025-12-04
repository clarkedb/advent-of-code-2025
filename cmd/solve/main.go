package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/clarkedb/advent-of-code-2025/solutions"
)

func main() {
	day := flag.Int("day", 0, "Day of the challenge (1-25)")
	part := flag.Int("part", 0, "Part of the challenge (1 or 2)")
	inputPath := flag.String("input", "", "Input file (defaults to input/day{N}.txt)")
	flag.Parse()

	if *day < 1 || *day > 25 {
		fmt.Fprintln(os.Stderr, "Error: day must be between 1 and 25")
		os.Exit(1)
	}

	if *part != 1 && *part != 2 {
		fmt.Fprintln(os.Stderr, "Error: part must be 1 or 2")
		os.Exit(1)
	}

	if *inputPath == "" {
		*inputPath = fmt.Sprintf("input/day%02d.txt", *day)
	}

	input, err := os.Open(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	solver, ok := solutions.Get(*day, *part)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: solution for day %d part %d not found\n", *day, *part)
		os.Exit(1)
	}

	start := time.Now()
	result, err := solver(input)
	duration := time.Since(start)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running solution: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("day %d - part %d [completed in %.4f ms]\n", *day, *part, float64(duration.Microseconds())/1000)
	fmt.Printf(">> %s\n", result)
}
