package main

import "fmt"

func main() {
	// declare with type
	var integer int
	// It receive the default value of the type
	fmt.Println(`The value of integer is`, integer)

	// declare receiving type inference
	var integerReceived = 10
	fmt.Println("The value of integer is", integerReceived)

	// Short hand declaration
	floatPoint := 3.4
	fmt.Println("The value of integer is", floatPoint)
}
