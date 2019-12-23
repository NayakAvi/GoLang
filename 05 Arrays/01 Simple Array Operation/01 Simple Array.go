package main

import "fmt"

func main() {

	// 1 Simple Declaration
	var a [5]int;
	a[0] = 1;
	a[1] = 2;
	a[2] = 3;
	a[3] = 4;
	a[4] = 5;
	fmt.Println("Array Elements a : " ,a)


	// 2 Declare and initialize
	var b = [5]int{1,2,3,4,5};
	fmt.Println("Array Elements b: " ,b)

	// 3 Declare and initialize
	var c = [5]int{1};
	fmt.Println("Array Elements c: " ,c)

	// 4 Pass by value
	var d = a
	d[0] = 100
	fmt.Println("Array Elements d: " ,d)

	//Array without Length, Compiler inteprets
	var e = [...]int{1,2,3,4,5};
	fmt.Println("Array Elements e: " ,e)

	
}
