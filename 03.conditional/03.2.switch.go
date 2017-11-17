package main

import "fmt"
import "time"

func main () {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("i don't know")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //or
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
}
