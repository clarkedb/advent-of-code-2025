package solutions

import "io"

// Solver is the function signature for all day/part solutions.
// It takes the input as an io.Reader, allowing solvers to read
// line-by-line with bufio.Scanner or read all with io.ReadAll.
type Solver func(input io.Reader) (string, error)

// registry holds all registered solutions indexed by day and part.
var registry = make(map[int]map[int]Solver)

// Register adds a solver function for a specific day and part.
func Register(day, part int, solver Solver) {
	if registry[day] == nil {
		registry[day] = make(map[int]Solver)
	}
	registry[day][part] = solver
}

// Get retrieves a solver function for a specific day and part.
func Get(day, part int) (Solver, bool) {
	if dayMap, ok := registry[day]; ok {
		solver, ok := dayMap[part]
		return solver, ok
	}
	return nil, false
}
