package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	// print last element
	fmt.Println(mySlice[(len(mySlice))-1])
	// print first element
	fmt.Println(mySlice[0])
}