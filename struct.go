package main

import (
	"fmt"
)

type Customer_bank_details struct {
	id            int
	Name          string
	Mobile_number int
	Address       string
}

func main() {
	var c1 Customer_bank_details = Customer_bank_details{1, "Aman", 1234567890, "Lko"}
	fmt.Println(c1.Name)

	var c2 Customer_bank_details = Customer_bank_details{2, "Raman", 15456565, "Delhi"}
	fmt.Println(c2.Name)

}
