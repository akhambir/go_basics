package main

import "fmt"

func main() {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over",
	    "the", "lazy", "dog"}                       // [] объявляет slice. динамически расширяемый
	fmt.Printf("%d\n", len(words))           // а так же передается в метод копия ссылки и может быть изменяемым

	printer2(words)
	printer2(words)

	words2 := make([]string, 4)                     // если нам нужно создать объект, но наполнять его потом
	words2[0] = "1"                                 // мы можем использовать функцию make()
	words2[1] = "2"
	words2[2] = "3"
	words2[3] = "4"

	printer2(words2);
	printer2(words2);
}

func printer2(words []string)  {
	for _, w := range words {
		fmt.Printf("%s\t", w)
	}

	fmt.Printf("\n")
	words[2] = "dog"
}
