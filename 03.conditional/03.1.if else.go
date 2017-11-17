package main

import "fmt"

func main() {
	// classic else if
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// multiple condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 { 
		fmt.Println(num, "is positive with one digit")
	} else {
		fmt.Println(num, "is positive with two digits")
	}
}