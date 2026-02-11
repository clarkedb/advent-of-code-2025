package solutions

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	Register(4, 1, day04Part1)
	Register(4, 2, day04Part2)
}

func day04Part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	grid := make([][]bool, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for i, c := range line {
			row[i] = (c == '@')
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	n := len(grid)
	m := len(grid[0])

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0

	for r := range n {
		for c := range m {
			if !grid[r][c] {
				continue
			}

			neighbors := 0
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < m && grid[nr][nc] {
					neighbors++
					if neighbors >= 4 {
						break
					}
				}
			}

			if neighbors < 4 {
				count++
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func day04Part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	grid := make([][]int, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			if c == '.' {
				row[i] = 1
			}
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	n := len(grid)
	m := len(grid[0])

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	prevCount := -1

	for iter := 1; count != prevCount; iter++ {
		prevCount = count
		for r := range n {
			for c := range m {
				if grid[r][c] != 0 && grid[r][c] <= iter {
					continue
				}

				neighbors := 0
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < n && nc >= 0 && nc < m && (grid[nr][nc] == 0 || grid[nr][nc] > iter) {
						neighbors++
						if neighbors >= 4 {
							break
						}
					}
				}

				if neighbors < 4 {
					count++
					grid[r][c] = iter + 1
				}
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}
