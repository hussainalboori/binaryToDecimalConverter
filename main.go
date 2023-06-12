package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){

	// arrgument len must be 3 exmple of arrgument "go run . In.text Out.text"
	if len(os.Args) != 3{
		fmt.Println("Error Reading The Arrgument ex. <go run . In.text Out.text>")
		return
	}
	in := os.Args[1]
	out := os.Args[2]

	//Read The File 

	InRead, err := ioutil.ReadFile(in)
	if err != nil {
		fmt.Println("\nError While Reading The File Please Check Again\n")
		return
	}

	// Check The File if It is Empty 

	if len(InRead) < 1 {
		fmt.Println("\nError Empty File Please Check Again\n")
		return
	}

	// Convert From Bin to Dec

	ConvertValue, err := BinaryToDecimal(string(InRead))
	if err != nil {
		fmt.Println("Error In Converting The file")
		return
	}
	// Write the Output 
	err = ioutil.WriteFile(out, []byte(fmt.Sprintf("%d\n", ConvertValue)), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
	fmt.Println("Done Successfully")
}
