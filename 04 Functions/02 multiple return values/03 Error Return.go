package main

import (
	"errors"
	"fmt"
)

func main() {
	value, error := devide(5, 0)
	if (error != nil) {
		fmt.Println("Error in devision function : ", error)
	} else {
		fmt.Println("Devisopn value is :", value)
	}
}

func devide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("devide by zero")
	} else {
		return x / y, nil
	}
}
