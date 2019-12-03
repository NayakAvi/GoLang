package main

import "fmt"

func main() {
	//constants can be typed or untyped
	const typedHello string = "Hello"


	var s string
	s = typedHello
	fmt.Println(s)

	var s2 = typedHello
	fmt.Println(s2)

	type MyString string
	var s3 MyString
	//s3 = typedHello Not ok

	const typedMyString MyString = "Hello"
	s3 = typedMyString

	fmt.Println(s3)

}
