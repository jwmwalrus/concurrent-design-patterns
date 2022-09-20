package main

import (
	"fmt"
	"time"

	"github.com/jwmwalrus/concurrent-design-patterns/futures/better/matrix"
)

func inverseProductSync(a, b *matrix.Type) *matrix.Type {
	aInv := matrix.Inverse(a)
	bInv := matrix.Inverse(b)
	return matrix.Product(aInv, bInv)
}

func inverseProductAsync(timeout time.Duration, a, b *matrix.Type) *matrix.Type {
	cancel := make(chan struct{})
	defer close(cancel)

	aInv := matrix.InverseAsync(cancel, a)
	bInv := matrix.InverseAsync(cancel, b)

	select {
	case <-time.After(timeout):
		cancel <- struct{}{}
		return nil
	case prod := <-matrix.ProductAsync(cancel, <-aInv, <-bInv):
		return prod
	}
}

func main() {
	a := matrix.New(1, 2, 3, 4)
	fmt.Printf("Matrix A:\n%v\n", a)

	b := matrix.New(5, 6, 7, 8)
	fmt.Printf("Matrix B:\n%v\n", b)

	start := time.Now()
	fmt.Printf("The (sync) product of the inverses is:\n%v", inverseProductSync(a, b))
	fmt.Printf("\tSync version took %v to complete\n\n", time.Since(start))

	for _, d := range []int{500, 1000, 1500, 2000} {
		timeout := time.Duration(d) * time.Millisecond
		fmt.Printf("The (async) product of the inverses (with timeout=%v) is:\n", timeout)
		start = time.Now()
		prod := inverseProductAsync(timeout, a, b)
		if prod != nil {
			fmt.Printf("%v", prod)
			fmt.Printf("\tAsync version took %v to complete\n\n", time.Since(start))
			continue
		}
		fmt.Printf("\tTimeout (%v) exceeded!\n\n", timeout)
	}
}
