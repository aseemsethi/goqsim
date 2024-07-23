// Learnings from https://github.com/takezo5096/goqkit
// and https://github.com/itsubaki/q/blob/main/q.go

package main

import (
	"fmt"
)

type Qsim struct {
	numQbits uint     // number of Qbits
	vec      VecQbits // Struct - having number of bits and their vector
	regQsim  []*Reg   // Not sure of this
	Name     string
}

func MakeCircuit(num uint, name string) Qsim {
	//var i uint
	//i = 1 << num
	reg := make([]*Reg, 0) // TBD - remove
	// Create a |0> vector for each of the input Qubits
	// and save in Circuit->vec
	fmt.Printf("Create a Qbit with ket0 - |0>\n")
	vec := NewVector(2) // 2 here means 2 dimension vector
	vec.Print()

	fmt.Printf("MakeCircuit with %d Qbits\n", num)
	return Qsim{numQbits: num, vec: vec, regQsim: reg, Name: name}

}

func main() {
	fmt.Println("Qsim starting...........................")
	// Create a circuit with n inputs, each input is |0>
	ckt := MakeCircuit(1, "Circuit1")
	fmt.Printf("Circuit %+v\n\n", ckt)
	fmt.Printf("Create Identity Matrix\n")
	iMatrix := NewIMatrix(2, 2)
	iMatrix.Print()

	fmt.Printf("Multiply a Vector with Identity Matrix\n")
	iMatrix.Print()
	ckt.vec.Print()
	v2 := iMatrix.Dot(ckt.vec)
	fmt.Printf("Result:\n")
	v2.Print()

}
