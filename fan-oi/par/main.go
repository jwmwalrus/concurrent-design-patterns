package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/generator/channel/primes"
)

const (
	many     = 1200
	nthreads = 4
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
	nums := primes.GetFirstN(many)

	var out []<-chan *rowType
	for i := 0; i < nthreads; i++ {
		out = append(out, fanOut(many/nthreads, nums))
	}

	in := fanIn(many, out...)

	for r := range in {
		i++
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

func fanOut(nsample int, in <-chan int) <-chan *rowType {
	out := make(chan *rowType, nsample)

	go func() {
		defer close(out)

		i := 0
		for n := range in {
			i++
			r := newRow(n)
			out <- r
			if i == nsample {
				break
			}
		}
	}()

	return out
}

func fanIn(ntotal int, out ...<-chan *rowType) <-chan *rowType {
	var wg sync.WaitGroup

	in := make(chan *rowType, ntotal)

	wg.Add(len(out))

	for _, och := range out {
		go func(rc <-chan *rowType) {
			defer wg.Done()

			for r := range rc {
				in <- r
			}
		}(och)
	}

	go func() {
		defer close(in)
		wg.Wait()
	}()

	return in
}
