package main

import (
	"fmt"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/generator/slice/primes"
)

func main() {
	start := time.Now()
	for i, p := range primes.GetFirstN(100) {
		fmt.Printf("%v ", p)
		if (i+1)%10 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\nGenerated in %v\n", time.Since(start))
}
