package main

import "fmt"

func main() {

	//type 1
	var i = 1
	for i = 1; i < 3; i++ {
		fmt.Println("Type1 : ", i)
	}

	//type 2
	i = 1
	for ; i < 3; i++ {
		fmt.Println("Type2 : ", i)
	}

	//type 3
	i = 1
	for ; i < 3; {
		fmt.Println("Type3 : ", i)
		i++
	}

	//type 4
	i = 1
	for {
		if (i > 2) {
			break
		}
		fmt.Println("Type4 : ", i)
		i++
	}

	//type 5
	for j := 1; ; {
		if (j > 2) {
			break
		}
		if (j < 2) {
			fmt.Println("Type5 : ", j)
			j++

		}

	}

}
