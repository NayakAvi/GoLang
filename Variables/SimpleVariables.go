package main

import "fmt"

func main() {
	//different types of declaring the variables

	/**
	types of variables supported

	1) bool
	2) string
	3) int, int8, int16
	4) float32, float64

	*/

	//type 1
	var user1FirstName string
	var user1SecondName string
	var user1Age int

	//type 2
	var user2FirstName, user2SecondName string
	var user2Age int

	//type 3

	var (
		user3FirstName  string
		user3SecondName string
		user3Age        int
	)

	user1FirstName = "user1FirstName"
	user1SecondName = "user1SecondName"
	user1Age = 10

	user2FirstName = "user2FirstName"
	user2SecondName = "user2SecondName"
	user1Age = 20

	user3FirstName = "user3FirstName"
	user3SecondName = "user3SecondName"
	user3Age = 10

	fmt.Println(user1FirstName)
	fmt.Println(user1SecondName)
	fmt.Println("user1Age : ", user1Age)

	fmt.Println(user2FirstName)
	fmt.Println(user2SecondName)
	fmt.Println("user2Age : ", user2Age)

	fmt.Println(user3FirstName)
	fmt.Println(user3SecondName)
	fmt.Println("user3Age : ", user3Age)

}
