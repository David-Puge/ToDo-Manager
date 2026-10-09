package task

import (
	"testing"
)

func TestAdd(t *testing.T) {

	var TestTasks []Task
	Add(&TestTasks, "Тренировка", "Присед, Жим, Тяга")

	if len(TestTasks) != 1 {
		t.Fatalf("Ожидалась длина слайса 1, но получилось %d ", len(TestTasks))
	}

	if TestTasks[0].Name != "Тренировка" {
		t.Fatalf("Ожидалось имя задачи - Тренировка -, но получилось - %s -", TestTasks[0].Name)
	}
	if TestTasks[0].Description != "Присед, Жим, Тяга" {
		t.Fatalf("Ожидалось описание задачи - Присед, Жим, Тяга -, но получилось - %s -", TestTasks[0].Description)
	}
	if TestTasks[0].ID != 1 {
		t.Fatalf("Ожидалось ID задачи - 1 -, но получилось - %d -", TestTasks[0].ID)
	}
	if TestTasks[0].Done != false {
		t.Fatalf("Ожидалось статус задачи - false -, но получилось - %t -", TestTasks[0].Done)
	}

	Add(&TestTasks, "Сходить в магазин", "Молоко, Курица, Рис")

	if len(TestTasks) != 2 {
		t.Fatalf("Ожидалась длина слайса 2, но получилось %d ", len(TestTasks))
	}

	if TestTasks[1].ID != 2 {
		t.Fatalf("Ожидалось ID задачи - 2 -, но получилось - %d -", TestTasks[0].ID)
	}

}

func TestDel(t *testing.T) {
	var TestTasks []Task
	Add(&TestTasks, "Тренировка", "Присед, Жим, Тяга")
	Add(&TestTasks, "Сходить в магазин", "Молоко, Курица, Рис")
	Add(&TestTasks, "Учеба GO", "Изучить работу с json")

	err := Del(&TestTasks, 2)
	if err != nil {
		t.Fatalf("Del вернул ошибку: %v", err)
	}

	if len(TestTasks) != 2 {
		t.Fatalf("Ожидалась длина слайса 2, но получилось %d ", len(TestTasks))
	}

	if TestTasks[0].ID != 1 || TestTasks[1].ID != 2 {
		t.Fatalf("Ожидалось ID задач 1 и 2 (упорядочить ID), но получилось %d и %d ", TestTasks[0].ID, TestTasks[1].ID)
	}

}

func TestSetDone(t *testing.T) {
	var TestTasks []Task
	Add(&TestTasks, "Тренировка", "Присед, Жим, Тяга")

	SetDone(&TestTasks, 1)
	if TestTasks[0].Done != true {
		t.Fatalf("Ожидалось статус задачи - true -, но получилось - %t -", TestTasks[0].Done)
	}
	SetDone(&TestTasks, 1)
	if TestTasks[0].Done != false {
		t.Fatalf("Ожидалось статус задачи - false -, но получилось - %t -", TestTasks[0].Done)
	}
}

func TestList(t *testing.T) {
	TestTasks := []Task{
		{ID: 1, Name: "Task 1", Description: "desc 1", Done: false},
		{ID: 2, Name: "Task 2", Description: "desc 2", Done: true},
		{ID: 3, Name: "Task 3", Description: "desc 3", Done: false},
		{ID: 4, Name: "Task 4", Description: "desc 4", Done: true},
		{ID: 5, Name: "Task 5", Description: "desc 5", Done: false},
	}

	all, err := List(&TestTasks, "all")
	if err != nil {
		t.Fatalf("Ошибка при выводе списка задач: %v", err)
	}

	if len(all) != 5 {
		t.Fatalf("Ожидалась длина слайса 5, но получилось %d ", len(all))
	}

	done, err := List(&TestTasks, "done")
	if err != nil {
		t.Fatalf("Ошибка при выводе списка задач: %v", err)
	}

	if len(done) != 2 {
		t.Fatalf("Ожидалась длина слайса 2, но получилось %d ", len(done))
	}

	pending, err := List(&TestTasks, "pending")
	if err != nil {
		t.Fatalf("Ошибка при выводе списка задач: %v", err)
	}

	if len(pending) != 3 {
		t.Fatalf("Ожидалась длина слайса 3, но получилось %d ", len(pending))
	}

}
