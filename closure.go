//Anonymous functions are useful when you want to define a function inline without having to name it.
//go supports anonymous func which can form closure
package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
