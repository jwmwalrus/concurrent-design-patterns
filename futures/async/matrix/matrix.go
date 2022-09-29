package matrix

import (
	"fmt"
	"time"
)

// Type declares a 2x2 matrix
type Type [2][2]float32

func (m *Type) String() string {
	return fmt.Sprintf("[[%v, %v]\n [%v, %v]]\n", m[0][0], m[0][1], m[1][0], m[1][1])
}

// New creates a matrix
func New(a, b, c, d int) *Type {
	return &Type{
		{float32(a), float32(b)},
		{float32(c), float32(d)},
	}
}

// Inverse inverts the given matrix
func Inverse(m *Type) *Type {
	return <-InverseAsync(m)
}

// Product returns the product of the given matrices
func Product(a, b *Type) *Type {
	return <-ProductAsync(a, b)
}

// InverseAsync inverts the given matrix, asynchronously
func InverseAsync(m *Type) chan *Type {
	doInverse := func(m *Type) *Type {
		time.Sleep(1 * time.Second)
		det := m[0][0]*m[1][1] - m[0][1]*m[1][0]
		return &Type{
			{m[1][1] / det, -1 * m[0][1] / det},
			{-1 * m[1][0] / det, m[0][0] / det},
		}
	}

	inv := make(chan *Type)

	go func() {
		defer close(inv)
		inv <- doInverse(m)
	}()
	return inv
}

// ProductAsync returns the product of the given matrices, asynchronously
func ProductAsync(a, b *Type) chan *Type {
	doProduct := func(a, b *Type) *Type {
		return &Type{
			{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
			{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
		}
	}

	prod := make(chan *Type)

	go func() {
		defer close(prod)
		prod <- doProduct(a, b)
	}()
	return prod
}
