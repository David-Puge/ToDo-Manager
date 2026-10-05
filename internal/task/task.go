package task

// Task представляет собой структуру задачи с полями ID, Name, Description и Done.
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
