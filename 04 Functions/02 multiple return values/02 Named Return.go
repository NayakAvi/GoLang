package main

func main() {
	sum, diff := calculateNamed(2, 2)
	println("Sum is : ", sum)
	println("Diff is : ", diff)
}

func calculateNamed(x int, y int) (sum, diff int) {
	sum = x + y // does not use := bz the named input declares the variable so we can just assign it
	diff = x - y
	return sum, diff
}