package solvers

// Solver defines the interface for day solvers
type Solver interface {
	Part1(input []byte) (string, error)
	Part2(input []byte) (string, error)
}

// DaySolverRegistry holds all registered day solvers
var DaySolverRegistry = make(map[int]Solver)

// RegisterSolver registers a solver for a specific day
func RegisterSolver(day int, solver Solver) {
	DaySolverRegistry[day] = solver
}

// GetSolver returns the solver for a given day, or nil if not registered
func GetSolver(day int) Solver {
	return DaySolverRegistry[day]
}

