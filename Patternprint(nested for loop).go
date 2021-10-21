
//1) triangle * pattern
package main

import (
	"fmt"
)

func main() {
	var n int = 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println("*")
		}
		fmt.Println(" ")

	}
}


//2)reverse triangle * pattern
package main

import (
	"fmt"
)

func main() {
	var n int = 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println("*")
		}
		fmt.Println(" ")

	}
}

//3)reverse triangle * pattern from right side
package main

import (
	"fmt"
)

func main() {
	var n int = 5
	for i := 1; i <= n; i++ {
		for j :=4 ; j >= i; j-- {
			fmt.Print(" ")
		}
        
        	for k :=1 ; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}

//4)complete  triangle * pattern 
package main

import (
	"fmt"
)

func main() {
	var n int = 5
	for i := 1; i <= n; i++ {
		for j :=4 ; j >= i; j-- {
			fmt.Print(" ")
		}
        
        	for k :=1 ; k <= i; k++ {
			fmt.Print(" *")
		}
		fmt.Println()

	}
}




