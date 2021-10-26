// in this program what we see is first it will print "defer start" after that loop will be executed in reversed order.
// that means for loop get executed already but it get stored in some memory and after that it get fetched in FIFO order.
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Defer start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)

	}

	fmt.Println("hello")
}
