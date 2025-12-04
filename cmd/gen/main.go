package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const solutionTemplate = `package solutions

import (
	"fmt"
	"strings"
)

func init() {
	Register({{.Day}}, 1, day{{.DayPadded}}Part1)
	Register({{.Day}}, 2, day{{.DayPadded}}Part2)
}

func day{{.DayPadded}}Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// TODO: implement solution
	return fmt.Sprintf("processed %d lines", len(lines)), nil
}

func day{{.DayPadded}}Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// TODO: implement solution
	return fmt.Sprintf("processed %d lines", len(lines)), nil
}
`

const testTemplate = `package solutions

import "testing"

func TestDay{{.DayPadded}}Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "line1\nline2",
			want:  "processed 2 lines",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day{{.DayPadded}}Part1(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day{{.DayPadded}}Part1() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDay{{.DayPadded}}Part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: "line1\nline2",
			want:  "processed 2 lines",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day{{.DayPadded}}Part2(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("day{{.DayPadded}}Part2() = %q, want %q", got, tt.want)
			}
		})
	}
}
`

type templateData struct {
	Day       int
	DayPadded string
}

func main() {
	day := flag.Int("day", 0, "Day to generate (1-25)")
	flag.Parse()

	if *day < 1 || *day > 25 {
		fmt.Fprintln(os.Stderr, "Error: day must be between 1 and 25")
		os.Exit(1)
	}

	data := templateData{
		Day:       *day,
		DayPadded: fmt.Sprintf("%02d", *day),
	}

	solutionPath := filepath.Join("solutions", fmt.Sprintf("day%02d.go", *day))
	testPath := filepath.Join("solutions", fmt.Sprintf("day%02d_test.go", *day))

	if err := generateFile(solutionPath, solutionTemplate, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating solution: %v\n", err)
		os.Exit(1)
	}

	if err := generateFile(testPath, testTemplate, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating test: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", solutionPath)
	fmt.Printf("Generated %s\n", testPath)
}

func generateFile(path, tmplContent string, data templateData) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("%s already exists", path)
	}

	tmpl, err := template.New("file").Parse(tmplContent)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
