package main

import (
	"os"
)

func main() {
	writer("Hello World!")
}

func writer(msg string) error {
	f, err := os.Create("helloworld.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(msg)
	return err
}
