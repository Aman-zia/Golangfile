// in some conditions when abnormal condition occurs in a program at that time we need to prematurily terminate this program then we use panic keyword.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 start")
	defer fmt.Println("2 defer")
	panic("cannot continue program")
	fmt.Println("3 panic")
}
