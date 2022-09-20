package main

import (
	"fmt"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/futures/async/matrix"
)

func inverseProductSync(a, b *matrix.Type) *matrix.Type {
	aInv := matrix.Inverse(a)
	bInv := matrix.Inverse(b)
	return matrix.Product(aInv, bInv)
}

func inverseProductAsync(a, b *matrix.Type) *matrix.Type {
	aInv := matrix.InverseAsync(a)
	bInv := matrix.InverseAsync(b)

	// NOTE: You can actually do some other work here

	return <-matrix.ProductAsync(<-aInv, <-bInv)
}

func main() {
	a := matrix.New(1, 2, 3, 4)
	fmt.Printf("Matrix A:\n%v\n", a)

	b := matrix.New(5, 6, 7, 8)
	fmt.Printf("Matrix B:\n%v\n", b)

	start := time.Now()
	fmt.Printf("The (sync) product of the inverses is:\n%v", inverseProductSync(a, b))
	fmt.Printf("\tSync version took %v to complete\n\n", time.Since(start))

	start = time.Now()
	fmt.Printf("The (async) product of the inverses is:\n%v", inverseProductAsync(a, b))
	fmt.Printf("\tAsync version took %v to complete\n", time.Since(start))
}
