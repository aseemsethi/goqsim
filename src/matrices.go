// The type complex128 in Golang is the set of all complex numbers with
// float64 real and imaginary parts.

package main

import (
	"fmt"
)

type VecQbits struct {
	N    uint
	Data []complex128
}

// Creates n Vectors each with a value of |0> and returns them in an array
func NewVector(n uint) VecQbits {
	var data []complex128
	var i uint

	fmt.Printf("New Vector...")

	for i = 0; i < n; i++ {
		data = append(data, complex128(0))
	}
	return VecQbits{N: n, Data: data}
}
