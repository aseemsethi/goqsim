// The type complex128 in Golang is the set of all complex numbers with
// float64 real and imaginary parts.

package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	Rows uint
	Cols uint
	Data [][]complex128
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

// CNOT Gate
// CNOT x 2Qbit - multiplies 4x4 CNTOT matrix to 4x1 vector
//    CNOT     2 Qbit  Vector representation
//  1 0 0 0 	[00]	[1 0 0 0]
//  0 1 0 0		[01]	[0 1 0 0]
//  0 0 0 1		[10]	[0 0 1 0]
//  0 0 1 0		[11]	[0 0 0 1]

func CNotGate() Matrix {
	var x, y uint
	m := NewMatrix(4, 4)
	// The following sets the diagnal to all 1s
	for x = 0; x < 4; x++ {
		for y = 0; y < 4; y++ {
			if x == y {
				m.Set(x, y, 1)
			} else {
				m.Set(x, y, 0)
			}
		}
	}
	// Now change the 2 values as per the CNOT matrix
	m.Set(2, 2, 0)
	m.Set(2, 3, 1)
	m.Set(3, 2, 1)
	m.Set(3, 3, 0)

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
