package main

import (
	"fmt"
)

func main() {
	// range is similar to range in python
	str_arr := []string{"a", "b", "c", "d"}
	for i, c := range str_arr {
		fmt.Println("Go", i)
		fmt.Println("kone", c)
	}

	//can also range over key/value pairs in maps:
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	for k, v := range m {
		fmt.Println("key:", k, "val:", v)
	}

	// can also just range over keys in maps:
	for k := range m {
		fmt.Println("key:", k)
	}
}
