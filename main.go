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
	res := algorithm.Solve(size, q)

	fmt.Println(*res)
}
