package main

import (
	"errors"
	"fmt"
)

func main() {
	value, error := devide1(5, 0)
	if (error != nil) {
		fmt.Println("Error in devision function : ", error)
	} else {
		fmt.Println("Devisopn value is :", value)
	}
}

//return fields are named
func devide1(x int, y int) (retValue int, retError error) {
	if y == 0 {
		retError = errors.New("devide by zero")
	} else {
		retValue = x / y
	}
	return
}
