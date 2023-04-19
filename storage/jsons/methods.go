package jsons

import (
	"bufio"
	"fmt"
	"os"
	"toDoList_app/storage/repo"
)


type AllTasks repo.AllTasks

func (t *AllTasks) GetAllTasks() []repo.Task {
	return t.Tasks
}

func (t *AllTasks) AddTask(task repo.Task) {

	reader := bufio.NewReader(os.Stdin)

	task.ID = len(t.Tasks) + 1
	var Title, Description, DueDate, Priority, StartDate string
	fmt.Println("Enter task title")
	Title, _ = reader.ReadString('\n')
	task.Title = Title[:len(Title)-1]
	fmt.Println("Enter task description")
	Description, _ = reader.ReadString('\n')
	task.Description = Description[:len(Description)-1]
	fmt.Println("Enter task priority")
	Priority, _ = reader.ReadString('\n')
	task.Priority = Priority[:len(Priority) - 1]
	fmt.Println("Enter task startDate")
	StartDate, _ = reader.ReadString('\n')
	task.StartDate = StartDate[:len(StartDate) - 1]
	fmt.Println("Enter task dueDate")
	DueDate, _ = reader.ReadString('\n')
	task.DueDate = DueDate[:len(DueDate) - 1]
	t.Tasks = append(t.Tasks, task)
}

func (t *AllTasks) RemoveTask(id int) {
	for i, val := range t.Tasks {
		if id == val.ID {
			t.Tasks = append(t.Tasks[:i],t.Tasks[i+1:]... )
		}
	}
}

func (t *AllTasks) MarkTasksCompleted(id int) {
	for i, val := range t.Tasks {
		if id == val.ID {
			t.Tasks[i].Completed = true
		}
	}
}