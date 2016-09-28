/* If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {
    limit := 1000
    var Sum int
    
    for x := 0; x < limit; x++ {
        if x%3 == 0 {
            //fmt.Println(x, "is divisible by 3")
            Sum += x
        } else if x%5 == 0 {
            //fmt.Println(x, "is divisible by 5")
            Sum += x
        }
    }
    fmt.Println("Sum is", Sum)
}
