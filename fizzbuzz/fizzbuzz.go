package main

import "fmt"

func main() {
	for i := 0; i <= 1000000; i++ {
		if i%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
