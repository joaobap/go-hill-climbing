package algorithm

import (
	"reflect"
)

// HillClimbingSolver is an interface that must be met to
// solve a problem using the Hill-Climbing algorithm.
type HillClimbingSolver interface {
	// New creates a new object of the problem to be solved.
	New(size int) HillClimbingSolver

	// Successor generates a new successor state from a problem object.
	// This function is responsable for the method used to generate the
	// successor, first choice or best choice.
	Successor() HillClimbingSolver

	// Objective checks if the current state is an objective state.
	Objective() bool
}

// HillClimbing is an implementation of the Hill Climbing algortihm
// for local search problems.
func HillClimbing(size int, h HillClimbingSolver) HillClimbingSolver {
	current := h.New(size)

	for {
		for i := 0; i < size*3; i++ {
			successor := current.Successor()

			// Found better successor
			if !reflect.DeepEqual(successor, current) {
				current = successor
			} else {
				break
			}
		}

		if current.Objective() {
			return current
		}
	}
}
