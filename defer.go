// defer is used before a calling function so it delays the execution of that function and that function returns the value at last.

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Defer start")
	student_registration()
}

func student_registration() {
	fmt.Println("Registration started")
	defer student_fees()
	fmt.Println("Registration ended")
}

func student_fees() {
	fmt.Println("Fees Paid")
}
