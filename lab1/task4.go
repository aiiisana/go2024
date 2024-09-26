package main

import "fmt"

func add(x int, y int) (sum int) {
	sum = x + y
	return
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func divide(x, y int) (int, int) {
	return x / y, x % y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(swap("hello", "world"))
	fmt.Println(divide(10, 5))
	q, w := divide(10, 3)
	fmt.Println(q, w)

}

// How do you define a function with multiple return values in Go?
// by listting them after parameters in ()

// What is the significance of named return values in Go?
//they act as a var inside the func and returned automatically

// How can you ignore certain return values if you don't need them?
// by using _
