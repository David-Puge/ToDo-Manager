package task

import (
	"fmt"
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
		return fmt.Errorf("Ошибка! ID не может быть меньше или ровен нулю./n")
	}
	if id > len(*tasks) {

		return fmt.Errorf("Ошибка! ID не может быть больше чем количество задач: n")
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
func List(tasks *[]Task, filter string) ([]Task, error) {

	resultTasks := make([]Task, 0)

	if filter == "done" {
		for _, task := range *tasks {
			if task.Done == true {
				resultTasks = append(resultTasks, task)
			}
		}
	} else if filter == "pending" {
		for _, task := range *tasks {
			if task.Done == false {
				resultTasks = append(resultTasks, task)
			}
		}
	} else if filter == "all" {
		for _, task := range *tasks {
			resultTasks = append(resultTasks, task)
		}
		return resultTasks, nil
	} else {
		return nil, fmt.Errorf("Ошибка! Фильтр может быть только all/done/pending")
	}
	return resultTasks, nil
}

func PrintTasks(tasks *[]Task) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "ID\tName\tDescription\tDone")
	for _, task := range *tasks {
		done := "[ ]"
		if task.Done {
			done = "[X]"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Name, task.Description, done)
	}
}

// SetDone изменяет статус задачи на противоположный (выполнено/не выполнено) по её ID.
func SetDone(tasks *[]Task, id int) error {
	if id <= 0 {
		return fmt.Errorf("Ошибка! ID не может быть меньше или ровен нулю./n")
	}
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Done = !(*tasks)[i].Done
		}
	}
	return nil
}
