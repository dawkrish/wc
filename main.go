package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("--------")

	var characterFlag bool
	var wordFlag bool
	var lineFlag bool
	var byteFlag bool

	flag.BoolVar(&characterFlag, "m", false, "Prints the number of characters")
	flag.BoolVar(&wordFlag, "w", false, "Prints the number of words")
	flag.BoolVar(&lineFlag, "l", false, "Prints the number of lines")
	flag.BoolVar(&byteFlag, "c", false, "Prints the number of bytes")

	flag.Parse()

	fmt.Println("characterFlag: ", characterFlag)
	fmt.Println("wordFlag: ", wordFlag)
	fmt.Println("lineFlag: ", lineFlag)
	fmt.Println("byteFlag: ", byteFlag)

	osArgs := flag.Args()
	// fmt.Println("Os Args -> ", osArgs)

	for _, fileName := range osArgs {
		_, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("%v : No such file exists\n", fileName)
		}
	}

	fmt.Println("--------")
}
