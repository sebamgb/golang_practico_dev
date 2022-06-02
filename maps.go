package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["a"] = 19
	fmt.Println(m1)
	fmt.Println(m1["a"])
}
