package main

import "fmt"

func main(){
	//if statement
	if x := 2; x > 0 {
		fmt.Println("positive")
	} else if x < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}

	//for loop
	sum := 0
	for i:=0; i<=10; i++{
		sum += i
	}
	fmt.Println("sum of first 10 int: ", sum)


	//switch
	day := 4

	switch day {
	case 1:
	  fmt.Println("Monday")
	case 2:
	  fmt.Println("Tuesday")
	case 3:
	  fmt.Println("Wednesday")
	case 4:
	  fmt.Println("Thursday")
	case 5:
	  fmt.Println("Friday")
	case 6:
	  fmt.Println("Saturday")
	case 7:
	  fmt.Println("Sunday")
	}
}


// How does the if statement in Go differ from other languages like Python or Java?
// can init var in if statement, no () for init

// What are the different ways to write a for loop in Go?
//classic     for i:=0; i<5; i++ {}
// while      for i < 5 {}
// infinite   for {}


// How does a switch statement in Go differ from switch in languages like C or Java?
// multiple expressions per case, can be used without expression like if series