package row

import (
	"fmt"
	"time"
)

// Type defines a row type
type Type struct {
	Number int
	Square int
	Cube   int
}

// New returns a new row
func New(n int) *Type {
	time.Sleep(100 * time.Microsecond)
	return &Type{n, n * n, n * n * n}
}

func (r *Type) String() string {
	return fmt.Sprintf("%v\t\t%v\t\t%v", r.Number, r.Square, r.Cube)
}
