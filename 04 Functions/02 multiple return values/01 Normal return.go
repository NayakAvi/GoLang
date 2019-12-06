package main

func main() {
	sum, diff := calculate(2, 2)
	println("Sum is : ", sum)
	println("Diff is : ", diff)
}

func calculate(x int, y int) (int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}
