// The type complex128 in Golang is the set of all complex numbers with
// float64 real and imaginary parts.

package main

import (
	"fmt"
)

type Matrix struct {
	Rows uint
	Cols uint
	Data [][]complex128
}

type VecQbits struct {
	N           uint
	Data        []complex128 // value of Qbit
	Probability []float64    // Probablity of each vector
}

// Creates n Vectors each with a value of |0> and returns them in an array
func NewVector(n uint) VecQbits {
	var data []complex128
	var i uint

	for i = 0; i < n; i++ {
		data = append(data, complex128(0))
	}
	return VecQbits{N: n, Data: data}
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

func NewIMatrix(r, c uint) Matrix {
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

func (m *Matrix) Print() {
	var i, j uint
	for i = 0; i < m.Rows; i++ {
		for j = 0; j < m.Cols; j++ {
			fmt.Printf("%.1f ", m.At(i, j))
		}
		fmt.Println("")
	}
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
