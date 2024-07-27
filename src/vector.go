package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	//"time"
)

type VecQbits struct {
	N    uint
	Data []complex128 // value of Qbit
	Prob []float64
}

// Creates n Vectors each with a value of |0> and returns them in an array
// Data[0] = 1 and Data[1] = 0 - making this a ket0 - |0>
// Each call to NewVector returns a single 2-dimension Qbit value of ket-0
func NewVector(dimension uint) VecQbits {
	var data []complex128
	//var i uint

	// for i = 0; i < n; i++ {
	// 	data = append(data, complex128(0))
	// }
	data = append(data, complex128(1))
	data = append(data, complex128(0))
	return VecQbits{N: dimension, Data: data}
}

// ************* Vector Utils
func (v *VecQbits) Set(n uint, val complex128) {
	v.Data[n] = val
}
func (v *VecQbits) SetAll(val complex128) {
	var i uint
	for i = 0; i < v.N; i++ {
		v.Set(i, val)
	}
}
func (v *VecQbits) At(n uint) complex128 {
	return v.Data[n]
}

// Probability returns the probability of q
// that is expressed as a Vector
func (v *VecQbits) Probability() []float64 {
	v.Prob = make([]float64, len(v.Data))
	for i, a := range v.Data {
		v.Prob[i] = math.Pow(cmplx.Abs(a), 2)
	}
	return v.Prob
}

func (v *VecQbits) Measure() int {
	if rand.Float64() < v.Prob[0] {
		// prob0
		return 0
	} else {
		//prob1
		return 1
	}
}

func (v *VecQbits) Print() {
	var i uint
	for i = 0; i < v.N; i++ {
		if i != 0 && i%16 == 0 {
			fmt.Println("")
			fmt.Printf("%d: ", i)
		}
		e := v.Data[i]
		fmt.Printf("")
		fmt.Printf("%.1f\n", e)
	}
	fmt.Printf("Probability: %.2f\n", v.Probability())
}
