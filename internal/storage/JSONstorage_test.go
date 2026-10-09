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
