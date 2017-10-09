package main

import (
	"fmt"
)

const (
	message = "Hello World!"
)

var (
	integer = 33
)

func main() {
	println(message)

	integer += 2
	print(integer)

	number := 15
	newMessage := "Hello everybody! %d\n"
	fmt.Printf(newMessage, number)


}
