package main

import "fmt"

func main() {
	//constants can be typed or untyped
	const typedHello = "Hello"


	var s string
	s = typedHello
	fmt.Println(s)

	var s2 = typedHello
	fmt.Println(s2)

	type MyString string
	var s3 MyString
	s3 = typedHello //OK

	fmt.Println(s3)

}
