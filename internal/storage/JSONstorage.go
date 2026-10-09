package storage

import (
	"ToDoManager/internal/task"
	"encoding/json"
	"fmt"
	"os"
)

const StoragePath = "todo.json"

// SaveJSON сохраняет список задач в формате JSON в основной файл хранения.
func SaveJSON(path string, tasks *[]task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("ошибка при кодировании JSON: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл файла %s: %w", path, err)
	}

	return nil
}

// LoadJSON загружает список задач из указаного файла в основной файла хранения в формате JSON.
func LoadJSON(JsonPath string, StoragePath string, tasks *[]task.Task) error {
	data, err := os.ReadFile(JsonPath)
	if err != nil {
		return fmt.Errorf("ошибка при проверке файла %s: %w", JsonPath, err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return fmt.Errorf("ошибка при переводе JSON %s: %w", JsonPath, err)
	}
	err = SaveJSON(StoragePath, tasks)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении JSON %s: %w", JsonPath, err)
	}

	return nil
}

// ExportJSON экспортирует список задач в формате JSON в указанный файл.
func ExportJSON(JsonPath string, tasks *[]task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("ошибка при кодировании JSON: %w", err)
	}

	_, err = os.Stat(JsonPath)
	if err != nil {
		return fmt.Errorf("ошибка при проверке файла %s: %w", JsonPath, err)
	}
	err = os.WriteFile(JsonPath, data, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл файла %s: %w", JsonPath, err)
	}

	return nil

}
