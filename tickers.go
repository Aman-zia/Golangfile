//tickers are used normally when you want to do something at regular interval of time.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello go learners")
	duration := time.Duration(10) * time.Second

	tk := time.NewTicker(duration)

	i := 0
	for range tk.C {
		i++
		Check()
		if i > 5 {
			tk.Stop()
			break
		}
	}
}

func Check() {
	fmt.Println("checking...........")
}
