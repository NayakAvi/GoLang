package main

import "fmt"

func main() {
	var num = 2

	//Scenario 1
	switch num {
	case 0:
		fmt.Println("Number is 0")
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Number is greater than 2")
	}



	//Scenario 2
	switch (num) { //with bracket
	case 0:
		fmt.Println("Number is 0")
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Number is greater than 2")
	}

	//Scenario 3
	switch { //with bracket
	case num==0:
		fmt.Println("Number is 0")
	case num==1:
		fmt.Println("Number is 1")
	case num==2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Number is greater than 2")
	}


	//Scenario 4
	switch newNum:=2; { //with bracket
	case newNum==0:
		fmt.Println("Number is 0")
	case newNum==1:
		fmt.Println("Number is 1")
	case newNum==2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Number is greater than 2")
	}

	//Scenario 5
	switch num {
	case 0,2,4,6,8:
		fmt.Println("Number is even")
	case 1,3,5,7,9:
		fmt.Println("Number is odd")
	default:
		fmt.Println("Number > 10")
	}








}
