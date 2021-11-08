// A recover can stop a panic from aborting the program and the execution continues

package main

import (
	"fmt"
)

func main() {

	// 6. Panics are not fatal
	// they are fatal only when we panic upto the go runtime
	// in which case, the go runtime kills this application

	fmt.Println("Main: Start")

	iCausePanics()
	fmt.Println("Main: Post Panic")
}

func iCausePanics() {

	fmt.Println("about to panic")

	panic("func: canot continue to execute")
	fmt.Println("i just panicked")
}

func tryToRecover() {

	if err := recover(); err != nil {
		fmt.Println("Error: ", err)
	}
}
