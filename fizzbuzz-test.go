package main

import "fmt"

// fizzbuzz function with integer input parameter
func fizzbuzz(input int) string {
	// if input is multiples of 3 and 5, return FizzBuzz
	if input%15 == 0 {
		return "FizzBuzz"
	}
	// if input is multiples of 3, return Fizz
	if input%3 == 0 {
		return "Fizz"
	}
	// if input is multiples of 5, return Buzz
	if input%5 == 0 {
		return "Buzz"
	}
	// if input is not multiples of 3 or 5, return the input number
	return fmt.Sprint(input)
}

func main() {
	// looping from 1 to 100
	for i := 1; i <= 100; i++ {
		// print the result from fizzbuzz function
		fmt.Println(fizzbuzz(i))
	}
}
