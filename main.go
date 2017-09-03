package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	size, _ := strconv.Atoi(os.Args[1])

	q := MakeQueen(size)

	resInterface := SolveWithHillClimbing(size, q)
	result := Queen(resInterface.(Queen))

	fmt.Printf("Final state:\n%v\n", result)
	fmt.Printf("Heuristic:\n%v\n", result.Heuristic())
}
