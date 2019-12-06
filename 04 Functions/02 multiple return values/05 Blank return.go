package main

import (
	"errors"
	"fmt"
)

func main() {
	// just checking for the error, hence the first parameter can be ignored using an _
	_, error := devide2(5, 0)
	if (error != nil) {
		fmt.Println("Error in devision function : ", error)
	}
}

//return fields are named
func devide2(x int, y int) (retValue int, retError error) {
	if y == 0 {
		retError = errors.New("devide by zero")
	} else {
		retValue = x / y
	}
	return
}
