package main

import "fmt"

func main() {

	// Shorthand declaration of array
	arr := [4]string{"facebook", "google", "amazon", "kloudone"}

	// Accessing the elements of

	fmt.Println("Elements of the array:")

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

}
