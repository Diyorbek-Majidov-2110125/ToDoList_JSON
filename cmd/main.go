package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"toDoList_app/storage/jsons"
	"toDoList_app/storage/repo"
)


func main() {
	var Tasks jsons.AllTasks
	
	data, err := ioutil.ReadFile("../storage/jsons/task.json")
	if err != nil {
		fmt.Println("Error reading:",err)
		return
	}
	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		fmt.Println("Error in unmarshalling:",err)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Println("------------------------------------------ToDo List Tasks------------------------------------------")
	fmt.Println("---------------------------------------------------------------------------------------------------")

	var response int
	for response != 5 {
		fmt.Println("---------------------------------------------------------------------------------------------------")
		fmt.Println("----------------------------------------------Choices:---------------------------------------------")
		fmt.Println("---------------------------------------------------------------------------------------------------\n")
		
		fmt.Println("<1> - See all tasks")
		fmt.Println("<2> - Add task to the list")
		fmt.Println("<3> - Remove task from the list")
		fmt.Println("<4> - Mark the task(completed)")
		fmt.Println("<5> - Exit")
		fmt.Println("\n---------------------------------------------------------------------------------------------------\n")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&response)

		var listofTasks []repo.Task
		switch response {
			case 1:
				listofTasks = Tasks.GetAllTasks()
				for _, t := range listofTasks {
					fmt.Println("--------------------------------------------------------------------------------------------------------------------")
fmt.Printf(`
Id: %v
title: %v
description: %v
completed: %v
priority: %v
startDate: %v
dueDate: %v

`,t.ID,t.Title,t.Description,t.Completed,t.Priority,t.StartDate,t.DueDate)
				}
			case 2:
				var task repo.Task
				Tasks.AddTask(task)
			case 3:
				var id int
				fmt.Println("Enter the task ID:")
				fmt.Scanln(&id)
				Tasks.RemoveTask(id)
			case 4:
				var id int
				fmt.Println("Enter the task ID:")
				fmt.Scanln(&id)
				Tasks.MarkTasksCompleted(id)
			case 5:
				fmt.Println("\nDo your best:)\n")
			default:
				fmt.Println("---------------------------------------------------------------------------------------------------")
				fmt.Println("----------------------------------------Invalid input,try again!-----------------------------------")
				fmt.Println("---------------------------------------------------------------------------------------------------")
		}



	}

	data, err = json.Marshal(Tasks)
	if err != nil {
		fmt.Println("Error in Marshalling:",err)
		return
	}

	err = ioutil.WriteFile("../storage/jsons/task.json",data,0644)
}
