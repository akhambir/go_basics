package main

import (
	"fmt"
)

func main() {
	printer("Hello World!")
	printer2("Hello", " World!")
	printer3("Hello", " World!", 4)
	err := printer4("err");
	println(err)
	auto_return("Hello!")
}

func printer(msg string)  {             // void method
	fmt.Printf("%s\n", msg)
}

func printer2(msg, msg2 string) {
	fmt.Printf("%s", msg)
	fmt.Printf("%s\n", msg2)
}

func printer3(msg, msg2 string, repeat int) {
	for repeat > 0 {
		fmt.Printf("%s", msg)
		fmt.Printf("%s\n", msg2)
		repeat--
	}
}

func printer4(msg string) error {
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func printer5(msg string) (string, error) {  // два возвращаемых значения
	msg += "\n"
	_, err := fmt.Printf(msg)
	return msg, err
}

func auto_return(msg string) (e error) {    // когда нет новых переменных то использовать := нельзя,
	_, e = fmt.Printf("%s\n", msg)   // так как уже известны типы.
	return                                  // в данном примере значение возвращается автоматически
}                                           // потому что оно объявлено в месте указания возвращаемого значения

func vararg_example(msgs ...string)  {      // vararg
	for _, msg := range msgs {              // forEach
		fmt.Printf("%s\n", msg)
	}
}