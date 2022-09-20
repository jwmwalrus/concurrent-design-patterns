package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/fan-oi/row"
	"github.com/jwmwalrus/concurrent-design-patterns/generator/channel/primes"
)

const (
	many     = 1200
	nthreads = 4
)

func main() {
	fmt.Printf("#\t\tprime\t\tsquare\n")

	var i, aggrPrimes, aggrSquares, aggrCubes int

	start := time.Now()
	nums := primes.GetFirstN(many)

	var out []<-chan *row.Type
	for i := 0; i < nthreads; i++ {
		out = append(out, fanOut(nums))
	}

	in := fanIn(out...)

	for r := range in {
		i++
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

func fanOut(in <-chan int) <-chan *row.Type {
	out := make(chan *row.Type, many/nthreads)

	go func() {
		defer close(out)

		for n := range in {
			r := row.New(n)
			out <- r
		}
	}()

	return out
}

func fanIn(out ...<-chan *row.Type) <-chan *row.Type {
	var wg sync.WaitGroup

	in := make(chan *row.Type, many)

	wg.Add(len(out))

	output := func(rc <-chan *row.Type) {
		defer wg.Done()

		for r := range rc {
			in <- r
		}
	}

	for _, och := range out {
		go output(och)
	}

	go func() {
		defer close(in)
		wg.Wait()
	}()

	return in
}
