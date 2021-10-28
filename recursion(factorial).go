// when a function calling itself again and again then it is known as recursion

package main

import (
	"fmt"
)

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	var i int = 5
	fmt.Printf("the factorial is %d\n", factorial(i))

}
