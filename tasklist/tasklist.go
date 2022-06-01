package tasklist

import "fmt"

type task struct {
	name        string
	description string
	complete    bool
}

func (t *task) chekComplete() {
	t.complete = true
}

func (t *task) changeDescription(description string) {
	t.description = description
}

func (t *task) changeName(name string) {
	t.name = name
}

func Tasklist_main() {
	t := &task{
		name:        "completar curso",
		description: "completar curso de golang esta semana",
	}

	fmt.Println(t)

	fmt.Printf("%+v\n", t)

	t.chekComplete()
	t.changeName("finalizar curso")
	t.changeDescription("finalizar curso practico de golang cuanto antes")
	fmt.Printf("%+v\n", t)
}
