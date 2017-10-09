package main

import "fmt"

func main() {
	atoz := "the quick brown fox jump over the lazy dog\n"

	for {
		fmt.Printf(atoz)        // infinity loop
	}

	for i, r := range atoz {

	}

	for _, r := range atoz {    // _ дает аозможность игнорировать параметр

	}

	counter := 0                // типа while
	for counter < 10 {
		fmt.Printf(atoz)
		counter += 1
	}

	for i := 0; i < 10; i++ {   // можно и так
		fmt.Printf(atoz)
	}

	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {   // и так
		fmt.Printf("%d Hello World!", j)
	}

	for true {
		fmt.Printf("I Love You!")
	}
}
