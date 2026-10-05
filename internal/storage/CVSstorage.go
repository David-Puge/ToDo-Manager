package storage

import (
	"ToDoManager/internal/task"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ExportCVS экспортирует список задач в формате CSV в указанный файл.
func ExportCVS(CsvPath string, tasks *[]task.Task) error {
	file, err := os.OpenFile(CsvPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Comma = ';'

	header := []string{"ID", "Name", "Description", "Done"}
	err = writer.Write(header)
	if err != nil {
		return fmt.Errorf("Ошибка записи заголовка: %w", err)
	}

	for _, task := range *tasks {
		idStr := strconv.Itoa(task.ID)
		doneStr := strconv.FormatBool(task.Done)

		row := []string{idStr, task.Name, task.Description, doneStr}

		err = writer.Write(row)
		if err != nil {
			return fmt.Errorf("Ошибка записи CSV: %w", err)
		}

	}
	return nil
}

// LoadCSV загружает список задач из файла в формате CSV.
func LoadCSV(CsvPath string, StoragePath string, tasks *[]task.Task) error {
	file, err := os.OpenFile(CsvPath, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Comma = ';'

	data, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("ошибка парсинга CSV: %w", err)
	}

	LoadTasks := []task.Task{}

	for i := 1; i < len(data); i++ {
		row := data[i]

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return fmt.Errorf("строка %d: неверный формат ID: %w", i+1, err)
		}
		done, err := strconv.ParseBool(row[3])
		if err != nil {
			return fmt.Errorf("строка %d: неверный формат Done: %w", i+1, err)
		}

		task := task.Task{
			ID:          id,
			Name:        row[1],
			Description: row[2],
			Done:        done,
		}

		LoadTasks = append(LoadTasks, task)
	}
	SaveJSON(StoragePath, &LoadTasks)

	return nil
}
