# advent-of-code-2025

Solutions in Go for the [2025 Advent of Code](https://adventofcode.com/2025)

## Usage

### Generate scaffolding for a new day

```bash
make gen DAY=2
```

This creates `solutions/day02.go` and `solutions/day02_test.go` with boilerplate code.

### Run a solution

```bash
make run DAY=1 PART=1
```

Or build the binary and run it directly:

```bash
make build
./bin/aoc -day 1 -part 1
```

By default, input is read from `input/day{NN}.txt`. You can specify a custom input file:

```bash
./bin/aoc -day 1 -part 1 -input path/to/input.txt
```

### Run tests

```bash
make test
```
