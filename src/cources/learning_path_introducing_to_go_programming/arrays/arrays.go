package main

import "fmt"

func main() {
	words := [...]string{"the", "quick", "brown", "fox"}  // [...] массив не фиксированной длинны
	fmt.Printf("%s\n", words[2])                   // [n] массив размерностью n

	printer(words)                     // в Go в качестве параметра array передается его копия, а не ссылка
}

func printer(words [4]string)  {
	for _, w := range words {
		fmt.Printf("%s\t", w)
	}
}