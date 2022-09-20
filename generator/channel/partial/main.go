package main

import (
	"fmt"
	"runtime"

	"github.com/jwmwalrus/concurrent-design-patterns/generator/channel/primes"
)

func main() {
	primes := primes.GetFirstN(20)

	fmt.Printf("Number of goroutines:%v\n", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", <-primes)
	}
	fmt.Printf("\n")

	fmt.Printf("\nNumber of goroutines:%v\n", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", <-primes)
	}
	fmt.Printf("\n")

	fmt.Printf("\nNumber of goroutines:%v\n", runtime.NumGoroutine())
	for i := 0; i < 15; i++ {
		p, ok := <-primes
		if !ok {
			fmt.Printf("\nChannel is closed!")
			break
		}
		fmt.Printf("%v ", p)
	}
	fmt.Printf("\n")
}
