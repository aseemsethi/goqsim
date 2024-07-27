// Learnings from https://github.com/takezo5096/goqkit
// and https://github.com/itsubaki/q/blob/main/q.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Qsim struct {
	numQbits uint     // number of Qbits
	vec      VecQbits // Struct - having number of bits and their vector
	Name     string
}

func MakeCircuit(num uint, name string) Qsim {
	//var i uint
	//i = 1 << num
	// Create a |0> vector for each of the input Qubits
	// and save in Circuit->vec
	fmt.Printf("Create a Qbit with ket0 - |0>\n")
	vec := NewVector(2) // 2 here means 2 dimension vector
	vec.Print()

	fmt.Printf("MakeCircuit with %d Qbits\n", num)
	return Qsim{numQbits: num, vec: vec, Name: name}

}

func main() {
	fmt.Println("Qsim starting...........................")
	rand.Seed(time.Now().UnixNano())

	// Create a circuit with n inputs, each input is |0>
	ckt := MakeCircuit(1, "Circuit1")
	fmt.Printf("Circuit %+v\n\n", ckt)
	fmt.Printf("Measuring.....: %d\n", ckt.vec.Measure())

	fmt.Printf("\nCreate Identity Matrix\n")
	iMatrix := IDGate(2, 2)
	iMatrix.Print()
	fmt.Printf("Multiply Qbit Vector with Identity Matrix\n")
	v2 := iMatrix.Dot(ckt.vec)
	v2.Print()
	fmt.Printf("Measuring.....: %d\n", v2.Measure())

	fmt.Printf("\nCreate Not Gate\n")
	iMatrix = NotGate(2, 2)
	iMatrix.Print()
	fmt.Printf("Not Gate on Qbit\n")
	v2 = iMatrix.Dot(ckt.vec)
	v2.Print()
	fmt.Printf("Measuring.....: %d\n", v2.Measure())

	fmt.Printf("\nCreate Hadamard Gate\n")
	iMatrix = HadGate()
	iMatrix.Print()
	fmt.Printf("Had Gate on Qbit\n")
	v2 = iMatrix.Dot(ckt.vec)
	v2.Print()
	fmt.Printf("Measuring.....: %d\n", v2.Measure())

	fmt.Printf("\nCreate CNot Gate\n")
	iMatrix = CNotGate()
	iMatrix.Print()
	fmt.Printf("CNOT Gate on Qbit |11>\n")
	vec := NewVector(4) // 4 dimension vector
	//set vector to qbit |11>, which is [0001]
	vec.Set2QbitVal(3)
	v2 = iMatrix.Dot(vec)
	// We should get qbit |10> which is [0010]]
	vec.Print()
	v2.Print()

}
