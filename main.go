package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/errbap/goia/algorithm"
	"github.com/errbap/goia/nqueen"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	size, _ := strconv.Atoi(os.Args[1])

	q := nqueen.Make(size)

	resInterface := algorithm.HillClimbing(size, q)
	result := nqueen.Queen(resInterface.(nqueen.Queen))

	fmt.Printf("Final state:\n%v\n", result)
	fmt.Printf("Heuristic:\n%v\n", result.Heuristic())
}
