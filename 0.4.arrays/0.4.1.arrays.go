package main

import "fmt"

func main() {
	// empty array declaration
	var a [5]int
	fmt.Println("emp: ", a)

	// set array value
	a[4] = 100
	fmt.Println("set: ", a)
	// get array by index
	fmt.Println("get: ", a[4])
	// print array length
	fmt.Println("len: ", len(a))
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b)
	
	var twoD [2][3]int 
	for i := 0 ; i < 2 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			twoD[i][j] = i + j
		}
	} 
	fmt.Println("array 2d", twoD)

}