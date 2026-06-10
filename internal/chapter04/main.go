package main

import "fmt"

func main() {
	array := [5]int{1: 10, 2: 20}
	fmt.Println(array)

	array[2] = 35
	fmt.Println(array)

	arrayPointer := [5]*int{0: new(int), 1: new(int)}
	*arrayPointer[0] = 10
	*arrayPointer[1] = 20
	fmt.Println(arrayPointer)

	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array1 = array2
	fmt.Println(array1)

	var array1Pointer [3]*string
	array2Pointer := [3]*string{new(string), new(string), new(string)}
	*array2Pointer[0] = "Red"
	*array2Pointer[1] = "Blue"
	*array2Pointer[2] = "Green"

	array1Pointer = array2Pointer
	fmt.Println(array1Pointer)

	var array42 [4][2]int
	array42_2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	array42_3 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	array42_4 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array42)
	fmt.Println(array42_2)
	fmt.Println(array42_3)
	fmt.Println(array42_4)

	array42[0][0] = 10
	array42[0][1] = 20
	array42[1][0] = 30
	array42[1][1] = 40
	fmt.Println(array42)

	var array22 [2][2]int
	var array22_2 [2][2]int
	array22_2[0][0] = 10
	array22_2[0][1] = 20
	array22_2[1][0] = 30
	array22_2[1][1] = 40
	fmt.Println(array22_2)
	array22 = array22_2
	fmt.Println(array22)
	var array3 [2]int = array22[1]
	fmt.Println(array3)
	var value int = array22[1][0]
	fmt.Println(value)

	var arrayArg [1e6]int
	foo(&arrayArg)
}

func foo(array *[1e6]int) {
	fmt.Println(1e6, len(array))
}
