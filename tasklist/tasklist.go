package tasklist

import "fmt"

type task struct {
	name        string
	description string
	complete    bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteToList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
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
	t1 := &task{
		name:        "completar curso",
		description: "completar curso de golang esta semana",
	}

	t2 := &task{
		name:        "completar curso",
		description: "completar curso de python esta semana",
	}

	t3 := &task{
		name:        "completar curso",
		description: "completar curso de java esta semana",
	}

	list := taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	list.addToList(t3)
	// for i := 0; i < len(list.tasks); i++ {
	// 	fmt.Println("index:", i, "descripcion:", list.tasks[i].description)
	// }
	// fmt.Println("---o---")
	// for i, task := range list.tasks {
	// 	fmt.Println("index:", i, "descripcion:", task.description)
	// }

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
