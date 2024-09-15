package main

import ("fmt"; "math")

type Shape interface {
	Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

func main(){
	c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}

    PrintArea(c)
    PrintArea(r)
}


// How do you define and implement an interface in Go?
// type nameOfInterface interface

// What is the role of interfaces in achieving polymorphism in Go?
// interfaces allows to define a set of methods that a type must implement to satisfy the interface

// How can you check if a type implements a certain interface?
// type assertion{
//     var _ Shape = (*Circle)(nil)
// } 
// or a 
// type switch{
//     func checkType(i interface{}) {
//         switch i.(type) {
//         case Shape:
//             fmt.Println("Implements Shape")
//         default:
//             fmt.Println("Does not implement Shape")
//         }
//     }
// }
