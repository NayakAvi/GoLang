package main

import "fmt"

func main()  {
	var int_var int = 5
	var float_var float32 = 5.6
	
	fmt.Println("int_var :",int_var)
	fmt.Println("float_var:",float_var)

	/*
	* CASE 1
	*/
	//explicitly cast
	i :=  float32(int_var) + float_var
	fmt.Println("i:",i)

	/*
	* CASE 2
	 */
	//explicitly cast
	var float_var2 float32 = float32(int_var)
	fmt.Println("float_var2:",float_var2)


}
