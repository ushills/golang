package main

import "fmt"

func main() {
    for x := 10; x >= 0; x-- {
		if x%2 != 0 {
			fmt.Println(x, "is odd")
		} else {
			fmt.Println(x, "is even")
		}
	}
}
