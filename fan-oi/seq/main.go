package main

import (
	"fmt"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/fan-oi/row"
	"github.com/jwmwalrus/concurrent-design-patterns/generator/channel/primes"
)

const (
	many = 1200
)

func main() {
	fmt.Printf("#\t\tprime\t\tsquare\n")

	var i, aggrPrimes, aggrSquares, aggrCubes int

	start := time.Now()
	for p := range primes.GetFirstN(many) {
		i++
		r := row.New(p)
		aggrPrimes += r.Number
		aggrSquares += r.Square
		aggrCubes += r.Cube
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
