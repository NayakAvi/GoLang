package main

import "fmt"

func main() {
	sum:=add(5,6)
	printSum(sum)
}

func printSum(sum int) {
	fmt.Println("Sum is :",sum)
}

func add(x int, y int) int{
	return x+y
}
