package repo

type ToDoList interface {
	GetAllTasks() []Task
	AddTask(task Task)
	RemoveTask(task Task)
	MarkTasksCompleted(index int)

}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    string `json:"priority"`
	StartDate   string `json:"startDate"`
}

type AllTasks struct{
	Tasks []Task
}