package main

import (
	"ToDoManager/internal/storage"
	"ToDoManager/internal/task"
	"encoding/json"
	"flag"
	"log"
	"os"
)

//===ToDoManager===
// 1. Основной функционал AddTask(Добавить задачу), DelTask(Удалить задачу), ListTask(Вывести список задач).
// Задачи можно сортировать по выполненным и не выполненным.
// Структура задачи (ID, Description, Done(bool))
// 2. Список задач должен храниться в todo.json и todo.CVS
// 3. Оформить команды и флаги в терминале
// 4. покрыть все тестами

// ===Команды===
// add -name "Название задачи" -desc "Описание задачи" - добавляет задачу в список
// del -id 1 - удаляет задачу по ID
// list -filter all/done/pending - выводит список задач в зависимости от фильтра
// complite -task 1 - меняет статус задачи на противоположный (выполнено/не выполнено) по её ID
// export -fail "Путь файла" -format json/csv - экспортирует список задач в указанный файл в формате JSON или CSV
// load -fail "Путь файла" -format json/csv - загружает список задач из указанного файла в формате JSON или CSV
func main() {
	Tasks := make([]task.Task, 0)
	if _, err := os.Stat(storage.StoragePath); os.IsNotExist(err) {
		file, err := os.Create(storage.StoragePath)

		if err != nil {
			log.Fatalf("Ошибка создания файла при инициализации: %v", err)
		}
		file.Close()
	}

	data, err := os.ReadFile(storage.StoragePath)
	if err != nil {
		log.Fatalf("Ошибка чтения файла при инициализации: %v", err)
	}

	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		log.Fatalf("Ошибка формирования json при инициализации: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Ошибка: Укажите команду")
	}

	switch os.Args[1] {
	case "add":
		addFlagSet := flag.NewFlagSet("add", flag.ExitOnError)
		TaskName := addFlagSet.String("name", " ", "Название новой задачи")
		TaskDescription := addFlagSet.String("desc", " ", "Описание новой задачи")
		err := addFlagSet.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("Ошибка парсинга add: %v", err)
		}

		task.Add(&Tasks, *TaskName, *TaskDescription)
		err = storage.SaveJSON(storage.StoragePath, &Tasks)
		if err != nil {
			log.Fatalf("Ошибка сохранения: %v", err)
		}

	case "delete":
		deleteFlagSet := flag.NewFlagSet("del", flag.ExitOnError)
		TasksID := deleteFlagSet.Int("id", 0, "id задачи для удаления")

		err := deleteFlagSet.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("Ошибка парсинга del: %v", err)
		}
		err = task.Del(&Tasks, *TasksID)
		if err != nil {
			log.Fatalf("Ошибка при удалении: %v", err)
		}
		err = storage.SaveJSON(storage.StoragePath, &Tasks)
		if err != nil {
			log.Fatalf("Ошибка сохранения: %v", err)
		}

	case "list":
		lisFlagSet := flag.NewFlagSet("list", flag.ExitOnError)
		filter := lisFlagSet.String("filter", "all", "Фильтр вывода задач(all, done, pending)")

		err := lisFlagSet.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("Ошибка парсинга del: %v", err)
		}

		list, err := task.List(&Tasks, *filter)
		if err != nil {
			log.Fatalf("Ошибка при выводе списка задач: %v", err)
		}
		task.PrintTasks(&list)

	case "complete":
		completeFlagSet := flag.NewFlagSet("complete", flag.ExitOnError)
		TasksID := completeFlagSet.Int("id", 0, "Меняет статус задачи по ID")

		err = completeFlagSet.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("Ошибка при парсинге аргумента: %v", err)
		}

		err = task.SetDone(&Tasks, *TasksID)
		if err != nil {
			log.Fatalf("Ошибка при изменении статуса задачи: %v", err)
		}
		err = storage.SaveJSON(storage.StoragePath, &Tasks)
		if err != nil {
			log.Fatalf("Ошибка сохранения: %v", err)
		}

	case "export":
		exportFlagSet := flag.NewFlagSet("export", flag.ExitOnError)
		failPath := exportFlagSet.String("out", "", "Путь файла для экспорта")
		format := exportFlagSet.String("format", "", "Формат сохранения JSON или CSV")

		err = exportFlagSet.Parse(os.Args[2:])

		switch *format {
		case "json":
			err = storage.ExportJSON(*failPath, &Tasks)
			if err != nil {
				log.Fatalf("Ошибка сохранения: %v", err)
			}
		case "csv":
			err = storage.ExportCVS(*failPath, &Tasks)
			if err != nil {
				log.Fatalf("Ошибка экспорта: %v", err)
			}
		default:
			log.Fatal("Не известный формат.")
		}

	case "load":
		loadFlagSet := flag.NewFlagSet("load", flag.ExitOnError)
		failPath := loadFlagSet.String("file", "", "Путь файла для импорта")
		format := loadFlagSet.String("format", "", "Формат сохранения JSON или CSV")

		err = loadFlagSet.Parse(os.Args[2:])

		switch *format {
		case "json":
			err = storage.LoadJSON(*failPath, storage.StoragePath, &Tasks)
			if err != nil {
				log.Fatalf("Ошибка импорта: %v", err)
			}
		case "csv":
			err = storage.LoadCSV(*failPath, storage.StoragePath, &Tasks)
			if err != nil {
				log.Fatalf("Ошибка импорта: %v", err)
			}
		default:
			log.Fatal("Не известный формат.")
		}

	}
}
