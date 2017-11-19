package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//seter
	fmt.Println("set", s)
	// geter
	fmt.Println("get", s[2])
	// show array length
	fmt.Println("len", len(s))

	// append data to array
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append", s)

	// create new array as c with same length with s
	// copy s data to c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	// slicing index 2-5
	l := s[2:5]
	fmt.Println("slicing 2:5", l)

	l = s[2:]
	fmt.Println("slicing 2:...", l)

	l = s[:4]
	fmt.Println("slicing ...:4", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("declare:", t)

	// create new array 2d
	twoD := make([][]int, 3)
	fmt.Println("Declare 2d array", twoD)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Nextdata 2d array", twoD)
}
