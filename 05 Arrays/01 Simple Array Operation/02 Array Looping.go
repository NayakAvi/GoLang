package main

import "fmt"

func main() {
	//simple for looping
	var a [5]int;
	a[0] = 1;
	a[1] = 2;
	a[2] = 3;
	a[3] = 4;
	a[4] = 5;
	var sum int

	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println("Sum of Array Elements in simple loop: ", sum)


	// Range loop
	sum = 0
	for _,value := range a {
		sum +=value
	}
	fmt.Println("Sum of Array Elements in range query: ", sum)


}
