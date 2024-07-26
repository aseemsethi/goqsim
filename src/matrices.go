// The type complex128 in Golang is the set of all complex numbers with
// float64 real and imaginary parts.

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

type Matrix struct {
	Rows uint
	Cols uint
	Data [][]complex128
}

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

func NewMatrix(r, c uint) Matrix {
	var data = make([][]complex128, r)
	var i, j uint
	for i = 0; i < r; i++ {
		data[i] = make([]complex128, c)
		for j = 0; j < c; j++ {
			data[i][j] = complex128(0)
		}
	}
	return Matrix{Rows: r, Cols: c, Data: data}
}

// ************* Martix utils

func IDGate(r, c uint) Matrix {
	m := NewMatrix(r, c)
	var i, j uint
	for i = 0; i < m.Rows; i++ {
		for j = 0; j < m.Cols; j++ {
			if i == j {
				m.Set(i, j, 1)
			}
		}
	}
	return m
}

func NotGate(r, c uint) Matrix {
	m := NewMatrix(r, c)
	m.Set(0, 0, 0)
	m.Set(0, 1, 1)
	m.Set(1, 0, 1)
	m.Set(1, 1, 0)
	return m
}

func YGate(r, c uint) Matrix {
	m := NewMatrix(r, c)
	m.Set(0, 0, complex(0, 0))
	m.Set(0, 1, complex(0, -1))
	m.Set(1, 0, complex(0, 1))
	m.Set(1, 1, complex(0, 0))
	return m
}

func HadGate() Matrix {
	sqrt2 := 1.0 / complex(math.Sqrt(2), 0)
	m := NewMatrix(2, 2)
	m.Set(0, 0, sqrt2)
	m.Set(0, 1, sqrt2)
	m.Set(1, 0, sqrt2)
	m.Set(1, 1, -sqrt2)
	return m
}

func (m *Matrix) Set(r, c uint, val complex128) {
	m.Data[r][c] = val
}

func (m *Matrix) SetAll(val complex128) {
	var i, j uint
	for i = 0; i < m.Rows; i++ {
		for j = 0; j < m.Cols; j++ {
			m.Set(i, j, val)
		}
	}
}

func (m *Matrix) At(r, c uint) complex128 {
	return m.Data[r][c]
}

// Matrix multiplication is only valid if the number of columns
// of the first matrix are equal to the number of rows of the
// second matrix; further, the resulting matrix will have the
// number of rows of the first matrix and the number of columns
// of the second matrix.
func (m *Matrix) Dot(v VecQbits) VecQbits {
	y := NewVector(m.Rows)
	var i, j uint
	for i = 0; i < m.Rows; i++ {
		var tmp complex128 = 0
		for j = 0; j < m.Cols; j++ {
			tmp += m.At(i, j) * v.At(j)
		}
		// Original QBit is not modified here.
		y.Set(i, tmp)
	}
	return y
}

func (m *Matrix) Print() {
	var i, j uint
	for i = 0; i < m.Rows; i++ {
		for j = 0; j < m.Cols; j++ {
			fmt.Printf("%.1f ", m.At(i, j))
		}
		fmt.Println("")
	}
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
	fmt.Printf("Probability: %.2f\n\n", v.Probability())
}
