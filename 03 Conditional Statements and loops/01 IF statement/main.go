package main

import "fmt"

func main() {
	/**
	IF Statements
	 */
	if(5>4){
		fmt.Println("Type1: 5 is greater than 4")
	}

	if 5>4 {
		fmt.Println("Type2: 5 is greater than 4")
	}

	if x:=5; x>4 {
		fmt.Println("Type3: 5 is greater than 4")
	}

	if x1,y1:=5,4; x1>y1{
		fmt.Println("Type4: 5 is greater than 4")
	}

	/**
	IF ELSE Statements
	*/

	if (4>5) {
		fmt.Println("Type5: 4 is greater than 5")
	} else{
		fmt.Println("Type5: 5 is greater than 4")
	}

	if (4>5) {
		fmt.Println("Type6: 4 is greater than 5")
	} else if 5>4  {
		fmt.Println("Type6: 5 is greater than 4")
	}

}
