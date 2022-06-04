package main

import "fmt"

type animal interface {
	move() string
}

func (dog) move() string {
	return "Soy un perro y estoy caminando"
}

func (fish) move() string {
	return "Soy un pez y estoy nadando"
}

func (bird) move() string {
	return "Soy un pajaro y estoy volando"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func Interfaces_main() {
	d := dog{}
	moveAnimal(d)
	f := fish{}
	moveAnimal(f)
	b := bird{}
	moveAnimal(b)
}
