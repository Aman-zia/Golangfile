//Regex is used for string matching.

package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "Working in Kloudone as a software engineer trainee"

	match1, err := regexp.MatchString("one", str)
	fmt.Println("Match: ", match1, " Error: ", err)

	str2 := "Computer Application"
	match2, err := regexp.MatchString("App", str2)
	fmt.Println("Match: ", match2, "Error: ", err)

	match3, err := regexp.MatchString("Chennai", str2)
	fmt.Println("Match: ", match3, "Error: ", err)
}
