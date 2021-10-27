//Pointer is a variable that is used to store the memory address of other variable

package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var ip *int
	ip = &x
	fmt.Printf("address of a variable %x\n", &x)
	fmt.Printf("address stored in ip variable %x\n", ip)
	fmt.Printf("value of ip variable %d\n", *ip)

}
