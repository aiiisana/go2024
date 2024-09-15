package main

import "fmt"

type Employee struct {
	Name string
	ID int
}

func (e Employee) Work() {
    fmt.Printf("Employee Name: %s, ID: %d\n", e.Name, e.ID)
}

type Manager struct {
    Employee
    Department string
}

func main(){
	m := Manager{
        Employee: Employee{
            Name: "Jhon",
            ID:   1234,
        },
        Department: "Engineering",
    }

    m.Work()
}

// What is embedding in Go, and how does it relate to composition?
// Embedding in go includes one struct within another, inheriting its fields and methods and composition is a way to build complex types using simpler ones.

// How does Go handle method calls on embedded types?
// they can be called directly on the outer struct

// Can an embedded type override a method from the outer struct?
// no, outer struct can define its own methods with the same name, but they are not overrides.