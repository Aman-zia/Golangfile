// channels are like pipes that connect concurrent goroutines

package main

import (
	"fmt"
)

func foo(c chan int, x int) {
	c <- x * 5
}
func main() {
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	v1 := <-fooVal
	v2 := <-fooVal

	fmt.Println(v1, v2)
}
