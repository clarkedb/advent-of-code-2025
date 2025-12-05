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

const startingPostition int = 50
const dialSize int = 100

func day01Part1(input io.Reader) (string, error) {
	var delta int
	var zeros int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			return "", fmt.Errorf("invalid input line: %q", line)
		}

		direction := rune(line[0])

		var value int
		_, err := fmt.Sscanf(line[1:], "%d", &value)
		if err != nil {
			return "", fmt.Errorf("invalid input line: %q", line)
		}

		switch direction {
		case 'R':
			// turn right
			delta += value
		case 'L':
			// turn left
			delta -= value
		default:
			return "", fmt.Errorf("invalid direction: %q", string(direction))
		}

		if (startingPostition+delta)%dialSize == 0 {
			zeros++
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", zeros), nil
}

func day01Part2(input io.Reader) (string, error) {
	var position int = startingPostition
	var zeros int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			return "", fmt.Errorf("invalid input line: %q", line)
		}

		direction := rune(line[0])

		var value int
		_, err := fmt.Sscanf(line[1:], "%d", &value)
		if err != nil {
			return "", fmt.Errorf("invalid input line: %q", line)
		}

		var crossings int
		switch direction {
		case 'R':
			// moving right: touch 0 each time we wrap from 99->0 or land on 0
			crossings = (position + value) / dialSize
			position = (position + value) % dialSize
		case 'L':
			// moving left: first touch 0 after 'position' steps (if position > 0),
			// then every dialSize steps after that
			if position == 0 {
				crossings = value / dialSize
			} else if position <= value {
				crossings = 1 + (value-position)/dialSize
			}
			position = (position - value%dialSize + dialSize) % dialSize
		default:
			return "", fmt.Errorf("invalid direction: %q", string(direction))
		}

		zeros += crossings
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", zeros), nil
}
