package main

import "fmt"

func main() {
	// var name type = expression
	// name := expression
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic looping
	for j := 7 ; j <= 9 ; j++ {
		fmt.Println(j)
	} 

	// break the loop
	for {
        fmt.Println("loop")
        break
	}
	
	for n := 0; n <= 5;n++ {
		if n%2 == 0 {
			// continue next iteration with continue
			continue
		}
		fmt.Println(n)
	}
}