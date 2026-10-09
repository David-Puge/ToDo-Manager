package storage

import (
	"ToDoManager/internal/task"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveJSON(t *testing.T) {
	TestTasks := []task.Task{
		{ID: 1, Name: "Тренировка", Description: "Присед, Жим, Тяга", Done: false},
		{ID: 2, Name: "Сходить в магазин", Description: "Молоко, Курица, Рис", Done: true},
	}
	filePath := filepath.Join(t.TempDir(), "test.json")
	err := SaveJSON(filePath, &TestTasks)
	if err != nil {
		t.Fatalf("SaveJSON вернул ошибку: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Ошибка чтения тестового файла json: %v", err)
	}
	getTask := []task.Task{}
	err = json.Unmarshal(data, &getTask)
	if err != nil {
		t.Fatalf("Ошибка парсинга тестового файла json: %v", err)
	}
	if len(getTask) != len(TestTasks) {
		t.Fatalf("Количество задач не совпадает")
	}
	if TestTasks[0].ID != 1 {
		t.Fatalf("Неверный ID. Ожидалось - 1 -, получили - %d - ", TestTasks[0].ID)
	}
	if TestTasks[0].Name != "Тренировка" {
		t.Fatalf("Неверное Имя. Ожидалось - Тренировка -, получили - %s - ", TestTasks[0].Name)
	}
	if TestTasks[0].Description != "Присед, Жим, Тяга" {
		t.Fatalf("Неверное Описание. Ожидалось - Присед, Жим, Тяга -, получили - %s - ", TestTasks[0].Description)
	}
	if TestTasks[0].Done != false {
		t.Fatalf("Неверный статус. Ожидалось - false -, получили - %t - ", TestTasks[0].Done)
	}

	if TestTasks[1].ID != 2 {
		t.Fatalf("Неверный ID. Ожидалось - 2 -, получили - %d - ", TestTasks[1].ID)
	}
	if TestTasks[1].Name != "Сходить в магазин" {
		t.Fatalf("Неверное Имя. Ожидалось - Сходить в магазин -, получили - %s - ", TestTasks[1].Name)
	}
	if TestTasks[1].Description != "Молоко, Курица, Рис" {
		t.Fatalf("Неверное Описание. Ожидалось - Молоко, Курица, Рис -, получили - %s - ", TestTasks[1].Description)
	}
	if TestTasks[1].Done != true {
		t.Fatalf("Неверный статус. Ожидалось - true -, получили - %t - ", TestTasks[1].Done)
	}

}

func TestLoadJSON(t *testing.T) {
	TestTasks := []task.Task{
		{ID: 1, Name: "Тренировка", Description: "Присед, Жим, Тяга", Done: false},
		{ID: 2, Name: "Сходить в магазин", Description: "Молоко, Курица, Рис", Done: true},
		{ID: 3, Name: "Практика в GO", Description: "Тестирование", Done: false},
		{ID: 4, Name: "Посмотреть лекции", Description: "ютуб", Done: true},
	}

	taskFilepath := filepath.Join(t.TempDir(), "test.json")
	taskFile, err := os.Create(taskFilepath)
	if err != nil {
		t.Fatalf("Ошибка создания тестового файла json: %v", err)
	}
	SaveJSON(taskFilepath, &TestTasks)
	taskFile.Close()

	emptyFilepath := filepath.Join(t.TempDir(), "testEmpty.json")
	emptyFile, err := os.Create(emptyFilepath)
	if err != nil {
		t.Fatalf("Ошибка создания тестового файла json: %v", err)
	}

	err = LoadJSON(taskFilepath, emptyFilepath, &TestTasks)
	if err != nil {
		t.Fatalf("Ошибка загрузки: %v", err)
	}

	data, err := os.ReadFile(emptyFilepath)
	if err != nil {
		t.Fatalf("Ошибка чтения тестового файла json: %v", err)
	}
	getTask := []task.Task{}
	err = json.Unmarshal(data, &getTask)
	if err != nil {
		t.Fatalf("Ошибка парсинга тестового файла json: %v", err)
	}

	if len(getTask) != len(TestTasks) {
		t.Fatalf("Количество задач не совпадает")
	}

	emptyFile.Close()
}

func TestExportJSON(t *testing.T) {
	TestTasks := []task.Task{
		{ID: 1, Name: "Тренировка", Description: "Присед, Жим, Тяга", Done: false},
		{ID: 2, Name: "Сходить в магазин", Description: "Молоко, Курица, Рис", Done: true},
	}
	filePath := filepath.Join(t.TempDir(), "test.json")
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Ошибка создания тестового файла json: %v", err)
	}
	file.Close()

	err = ExportJSON(filePath, &TestTasks)
	if err != nil {
		t.Fatalf("Ошибка экспорта: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Ошибка чтения тестового файла json: %v", err)
	}
	getTask := []task.Task{}
	err = json.Unmarshal(data, &getTask)
	if err != nil {
		t.Fatalf("Ошибка парсинга тестового файла json: %v", err)
	}
	if len(getTask) != len(TestTasks) {
		t.Fatalf("Количество задач не совпадает")
	}
}
