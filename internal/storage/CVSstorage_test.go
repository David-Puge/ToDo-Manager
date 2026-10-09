package storage

import (
	"ToDoManager/internal/task"
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExportCSV(t *testing.T) {
	TestTasks := []task.Task{
		{ID: 1, Name: "Тренировка", Description: "Присед, Жим, Тяга", Done: false},
		{ID: 2, Name: "Сходить в магазин", Description: "Молоко, Курица, Рис", Done: true},
		{ID: 3, Name: "Практика в GO", Description: "Тестирование", Done: false},
		{ID: 4, Name: "Посмотреть лекции", Description: "ютуб", Done: true},
	}

	filePath := filepath.Join(t.TempDir(), "testCSV.csv")
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Ошибка создания тестового файла csv: %v", err)
	}
	defer file.Close()

	err = ExportCVS(filePath, &TestTasks)
	if err != nil {
		t.Fatalf("Ошибка экспорта: %v", err)
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'

	data, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Ошибка чтения тестового файла csv: %v", err)
	}

	if len(data) != len(TestTasks)+1 {
		t.Fatalf("Количество строк в CSV не совпадает с количеством задач")
	}

}

func TestLoadCSV(t *testing.T) {
	TestTasks := []task.Task{
		{ID: 1, Name: "Тренировка", Description: "Присед, Жим, Тяга", Done: false},
		{ID: 2, Name: "Сходить в магазин", Description: "Молоко, Курица, Рис", Done: true},
		{ID: 3, Name: "Практика в GO", Description: "Тестирование", Done: false},
		{ID: 4, Name: "Посмотреть лекции", Description: "ютуб", Done: true},
	}
	filePath := filepath.Join(t.TempDir(), "testLoad.csv")

	err := ExportCVS(filePath, &TestTasks)
	if err != nil {
		t.Fatalf("Ошибка экспорта: %v", err)
	}
	emptyFilePath := filepath.Join(t.TempDir(), "testemptyCSV.csv")
	if err := os.WriteFile(emptyFilePath, []byte("[]"), 0644); err != nil {
		t.Fatalf("Ошибка создания файла для задач: %v", err)
	}

	err = LoadCSV(filePath, emptyFilePath, &TestTasks)
	if err != nil {
		t.Fatalf("Ошибка загрузки: %v", err)
	}

	GetTasks := []task.Task{}
	data, err := os.ReadFile(emptyFilePath)
	if err != nil {
		t.Fatalf("Ошибка при чтении файла: %v", err)
	}
	err = json.Unmarshal(data, &GetTasks)
	if err != nil {
		t.Fatalf("Ошибка парсинга тестового файла csv: %v", err)
	}

	if len(GetTasks) != len(TestTasks) {
		t.Fatalf("Ошибка, количество задач отличаются: %v", err)

	}

}

func TestLoadCSVRejectsIncompleteRow(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "incomplete.csv")
	csvData := "ID;Name;Description;Done\n1;Только имя\n"
	if err := os.WriteFile(filePath, []byte(csvData), 0644); err != nil {
		t.Fatalf("Ошибка создания тестового файла csv: %v", err)
	}

	err := LoadCSV(filePath, filepath.Join(t.TempDir(), "todo.json"), &[]task.Task{})
	if err == nil {
		t.Fatal("ожидалась ошибка для строки CSV с недостающими колонками")
	}
}
