package errba

type HillClimbingSolver interface {
	New(size int) HillClimbingSolver
	Successor() HillClimbingSolver
	Objective() bool
}

func Solve(size int, h HillClimbingSolver) *HillClimbingSolver {
	current := h.New(size)

	for {
		for i := 0; i < size*3; i++ {
			successor := current.Successor()

			if successor != nil {
				current = successor
				break
			} else {
				continue
			}
		}

		if current.Objective() {
			return &current
		}
	}
}
