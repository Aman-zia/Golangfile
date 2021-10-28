package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slc := arr[0:3]
	slc = append(slc, 45)
	fmt.Println(slc, len(slc), arr)
}
