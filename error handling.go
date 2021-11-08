package main

import (
	"errors"
	"fmt"
)

func division(num1, num2 int) (int, error) {
	if num2 == 0 {
		return -1, errors.New("Divide by zero")
	}
	return num1 / num2, nil
}

func main() {
	result, err := division(20, 10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("the Division of two numbers is:", result)
	}
}
