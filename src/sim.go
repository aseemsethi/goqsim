// Learnings from https://github.com/takezo5096/goqkit

package main

import (
	"fmt"
)

type Qsim struct {
	numQbits uint
	vec      VecQbits
	regQsim  []*Reg
	Name     string
}

func MakeCircuit(num uint, name string) Qsim {
	//var i uint
	//i = 1 << num
	reg := make([]*Reg, 0)
	// Create a |0> vector for each of the input Qubits
	// and save in Circuit->vec
	vec := NewVector(num)

	fmt.Printf("MakeCircuit with %d registers\n", num)
	return Qsim{numQbits: num, vec: vec, regQsim: reg, Name: name}

}

func main() {
	fmt.Println("Qsim start...")
	ckt := MakeCircuit(2, "Circuit1")
	fmt.Printf("Circuit %+v\n", ckt)

}
