package main

import (
	"fmt"
	"os"
)

func main(){
	in := os.Args[1]
	// out := os.Args[2]
	InRead, err := os.ReadFile(in)

	InReadString := string(InRead)

	if err != nil {
		fmt.Println("\nError While Reading The File Please Check Again\n")
	}

	if len(InRead) < 1 {
		fmt.Println("\nError Empty File Please Check Again\n")
	}
	fmt.Println(binaryToDecimal(InReadString))
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	binary := "101"
// 	decimal, err := binaryToDecimal(binary)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println(decimal)
// }
