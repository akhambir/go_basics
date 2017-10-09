package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello World!\n");

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output")
		fallthrough  // попадая сюда пройдем дальше в следующий кейс
	case n != 13:
		fmt.Printf("Wrong number of characters")
	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")
}
