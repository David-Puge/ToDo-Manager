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

	if TestTasks[0].ID != 1 && TestTasks[1].ID != 2 {
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
