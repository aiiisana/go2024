package main

import "fmt"

func main() { 
	var name string = "Aisana"
	var student = true
	var gpa float64 = 3.18	
	age := 18

	fmt.Printf("name: %s, type: %T\n", name, name)
	fmt.Printf("age: %d, type: %T\n", age, age)
	fmt.Printf("is student: %t, type: %T\n", student, student)
	fmt.Printf("gpa: %.2f, type: %T\n", gpa, gpa)
}

// Questions:
// What is the difference between using var and := to declare variables?
// using var you are specifying type or value for the variable,
// in case using := Go assigning the type and you must give value to the variable

// How do you print the type of a variable in Go?
// Printf("type: %T", variable)

// Can you change the type of a variable after it has been declared? Why or why not?
//no, because Go is statically typed, meaning the type of a variable is determined at compile-time, not runtime.