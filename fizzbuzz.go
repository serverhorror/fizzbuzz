package fizzbuzz

import (
	"fmt"
)

func FizzBuzz(fizzbuzzer *int) {
	for i := 1; i <= *fizzbuzzer; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// vim: set ts=4 sts=4 fenc=utf-8 noexpandtab list:
