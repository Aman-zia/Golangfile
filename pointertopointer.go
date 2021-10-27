package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var ip *int = &x
	var ipp **int = &ip

	fmt.Printf("The value of x is  %d\n", x)
	fmt.Printf("The value of x is  %d\n", *ip)
	fmt.Printf("The value of x is  %d\n", **ipp)

}
