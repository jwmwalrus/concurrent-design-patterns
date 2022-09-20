package main

import (
	"fmt"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/futures/sync/matrix"
)

func inverseProduct(a, b *matrix.Type) *matrix.Type {
	aInv := matrix.Inverse(a)
	bInv := matrix.Inverse(b)
	return matrix.Product(aInv, bInv)
}

func main() {
	a := matrix.New(1, 2, 3, 4)
	fmt.Printf("Matrix A:\n%v\n", a)

	b := matrix.New(5, 6, 7, 8)
	fmt.Printf("Matrix B:\n%v\n", b)

	start := time.Now()
	fmt.Printf("The product of the inverses is:\n%v", inverseProduct(a, b))
	fmt.Printf("\tIt took %v to complete\n", time.Since(start))
}
