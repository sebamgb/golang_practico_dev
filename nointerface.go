package main

import "fmt"

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) walk() {
	fmt.Println("Soy un perro y estoy caminando")
}

func (fish) swim() {
	fmt.Println("Soy un pez y estoy nadando")
}

func (bird) fly() {
	fmt.Println("Soy un pajaro y vuelo")
}

func moveDog(d dog) {
	d.walk()
}

func moveFish(f fish) {
	f.swim()
}

func moveBird(b bird) {
	b.fly()
}

func NoInterfaces_main() {
	d := dog{}
	moveDog(d)
	f := fish{}
	moveFish(f)
	b := bird{}
	moveBird(b)
}
