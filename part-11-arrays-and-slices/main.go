package main

import "fmt"

func setSection() {
	fmt.Println()
	fmt.Println("----------------------------")
	fmt.Println()
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func main() {
	fmt.Println("Arrays")
	// Declare an array
	var threeNillIntegers [3]int
	fmt.Println(threeNillIntegers)

	// Declare an not nil array
	var threeIntegers = [3]int{1, 2, 3}
	fmt.Println(threeIntegers)

	// Declare an auto-size array
	var integers = [...]int{1, 2, 3, 4, 5}
	fmt.Println(fmt.Sprintf(`Array %d. Size %d`, integers, len(integers)))

	setSection()

	// Arrays are passed by value, so it is 'cloned' when passed by parameter OR assigned to another variable
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	setSection()
	fmt.Println("Slices")

	// Slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	var sli []int = arr[1:4]
	fmt.Println(sli)

	// Slice from nothing
	slic := []int{6, 7, 8}
	fmt.Println(slic)

	// Every slice is a reference to an array. Values altered in slice will reflect in original array
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// Append slices together
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	// Slices are passed by reference, so changes inside the function will reflect on the slice outside. It is the opposite of an array.
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
}
