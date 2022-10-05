package main

import (
	"fmt"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/generator/channel/primes"
)

const (
	many = 1200
)

type rowType struct{ number, square, cube int }

func newRow(n int) *rowType {
	time.Sleep(100 * time.Microsecond)
	return &rowType{n, n * n, n * n * n}
}

func (r *rowType) String() string {
	return fmt.Sprintf("%v\t\t%v\t\t%v", r.number, r.square, r.cube)
}

func main() {
	fmt.Printf("#\t\tprime\t\tsquare\n")

	var i, aggrPrimes, aggrSquares, aggrCubes int

	start := time.Now()
	for p := range primes.GetFirstN(many) {
		i++
		r := newRow(p)
		aggrPrimes += r.number
		aggrSquares += r.square
		aggrCubes += r.cube
		fmt.Printf("%v\t\t%v\n", i, r)
	}
	fmt.Printf("\nIt took %v to generate\n", time.Since(start))
	fmt.Printf(
		"\nSum of primes: %v\nSum of squares: %v\nSum of cubes: %v\n",
		aggrPrimes,
		aggrSquares,
		aggrCubes,
	)
}
