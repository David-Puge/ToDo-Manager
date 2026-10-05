package task

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

// Add добавляет новую задачу в список задач.
func Add(tasks *[]Task, TaskName string, TaskDescription string) {

	var newTask Task
	newTask.ID = 1
	if len(*tasks) != 0 {
		newTask.ID = (*tasks)[len(*tasks)-1].ID + 1
	}
	newTask.Name = TaskName
	newTask.Description = TaskDescription

	newTask.Done = false

	*tasks = append(*tasks, newTask)
	fmt.Printf("Добавили, %s!, %s", newTask.Name, newTask.Description)

}

// Del удаляет задачу из списка задач по её ID.
func Del(tasks *[]Task, id int) error {
	if id <= 0 {
		err := fmt.Errorf("Ошибка! ID не может быть меньше или ровен нулю./n")
		return err
	}
	if id > len(*tasks) {
		err := fmt.Errorf("Ошибка! ID не может быть больше чем количество задач./n")
		return err
	}

	for i, task := range *tasks {

		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			for j := range *tasks {
				(*tasks)[j].ID = j + 1
			}
		}
	}
	return nil

}

// List выводит список задач в зависимости от фильтра.
func List(tasks *[]Task, filter string) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)
	defer w.Flush()

	if filter == "all" {
		fmt.Println("=== ВСЕ ЗАДАЧИ ===")
		for _, task := range *tasks {
			done := "[ ]"
			if task.Done == true {
				done = "[Х]"
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Name, task.Description, done)

		}
	}
	if filter == "done" {
		fmt.Println("=== ВЫПОЛНЕНЫЕ ЗАДАЧИ ===")
		for _, task := range *tasks {
			done := "[ ]"
			if task.Done == true {
				done = "[Х]"
			}
			if task.Done == true {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Name, task.Description, done)
			}
		}
	}
	if filter == "pending" {
		fmt.Println("=== АКТУАЛЬНЫЕ ЗАДАЧИ ===")
		for _, task := range *tasks {
			done := "[ ]"
			if task.Done == true {
				done = "[Х]"
			}
			if task.Done == false {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Name, task.Description, done)
			}
		}
	}
}

// SetDone изменяет статус задачи на противоположный (выполнено/не выполнено) по её ID.
func SetDone(tasks *[]Task, id int) {
	if id <= 0 {
		log.Fatal("Ошибка! ID не может быть меньше или ровен нулю./n")
		os.Exit(1)
	}
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Done = !(*tasks)[i].Done
		}
	}
}
