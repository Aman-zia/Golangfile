Go lang file

1) hello world program

package main

import "fmt" //it stands for format & it's a package

func main() {
	fmt.Println("Hello, world")
}

2) Adding two numbers 

package main

import "fmt"

func main() {
	num1 := 5 // var num1 = 5 or num1 :=5 both are equal
	var num2 = 5
	var result = num1 + num2
	fmt.Print(result)
}

3) how to find square root of a number and also floor ceil methods are used
package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 8
	var result = math.Sqrt(num)
	var result1 = math.Ceil(result)
	fmt.Printf("the output is %.2f thank u", result1)
}



4) Package level and function level
package main

import "fmt"

var a = 9 //package level

func demo() {
	a := 8 //function level
	fmt.Print(a)
}
func main() {
	demo()
	fmt.Print(a)
}




5) For loop
package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("kloudone")

	}

}




 6) odd even program taking input from the user

package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter a number :")
	var num int
	fmt.Scanf("%d", &num)

	if (num % 2) == 0 {
		fmt.Print("Number is even")
	} else {
		fmt.Print("number is odd")

	}

}





7) Palindrome number

package main

import "fmt"

func main() {
	var num, rem, temp int
	var rev int = 0

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&num)

	temp = num

	// For Loop used in format of While Loop
	for {
		rem = num % 10
		rev = rev*10 + rem
		num /= 10

		if num == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == rev {
		fmt.Printf("%d is a Palindrome", temp)
	} else {
		fmt.Printf("%d is not a Palindrome", temp)
	}

}

