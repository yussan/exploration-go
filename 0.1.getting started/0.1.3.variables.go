package main

import "fmt"

func main(){
	var a string = "initial"
	fmt.Println(a)

	// declare multiple variables
	var b, c int = 1, 2
	fmt.Println(b + c)

	// declare boolen
	var d = true
	fmt.Println(d && d)

	// declare variable and set value latter
	var e int 
	e = 12 
	fmt.Println(e)
}