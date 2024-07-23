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
	fmt.Printf("Create 2 Qbits and assign |0>\n")
	vec := NewVector(num)
	fmt.Printf("%+v\n\n", vec)

	fmt.Printf("MakeCircuit with these %d Qbits as inputs\n", num)
	return Qsim{numQbits: num, vec: vec, regQsim: reg, Name: name}

}

func main() {
	fmt.Println("Qsim starting...........................")
	// Create a circuit with n inputs, each input is |0>
	ckt := MakeCircuit(2, "Circuit1")
	fmt.Printf("Circuit %+v\n\n", ckt)
	fmt.Printf("Create Identity Matrix\n")
	iMatrix := NewIMatrix(2, 2)
	iMatrix.Print()

}
