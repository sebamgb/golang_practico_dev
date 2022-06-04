package main

import "fmt"

func Apuntadores_main() {
	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
}

// func changeValue(a int) {
// 	fmt.Println(&a)
// 	a = 40
// }
