package main

import (
	"fmt"
)

func main() {

	var g string = "A"

	switch {

	case g == "A":
		fmt.Println("good grades")

	case g == "B":
		fmt.Println("ok grades")

	case g == "C":
		fmt.Println("not so good grades")

	case g == "D":
		fmt.Println("fail grades")

	default:
		fmt.Println("please enter grade")
	}
}
