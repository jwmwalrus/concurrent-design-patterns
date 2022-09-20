package main

import (
	"fmt"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/generator/channel/primes"
)

func main() {
	start := time.Now()
	var i int
	for p := range primes.GetFirstN(100) {
		i++
		fmt.Printf("%v ", p)
		if i%10 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\nGenerated in %v\n", time.Since(start))
}
