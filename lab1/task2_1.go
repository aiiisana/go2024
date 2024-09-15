package main

import "fmt"

type Person struct {
    name string
    age  int
}

func newPerson(name string, age int) *Person {
    return &Person{
        name: name,
        age:  age,
    }
}

func (p Person) greet() {
    fmt.Println("Hi, I am", p.name, "! I am", p.age, "years old.")
}

func main() {
    p := newPerson("John", 24)
    p.greet()
}

// How do you define a struct in Go?
// by: type nameOfStruct struct {}

// How do methods differ from regular functions in Go?
// methods are associated with a struct, while func are not associated with any type.

// Can a method in Go be associated with types other than structs?
// no