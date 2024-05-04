package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const filename = "todo.txt"

func main() {
	fmt.Println("Welcome to Todo List!")

	// Load existing tasks
	tasks := loadTasks()

	// Menu loop
	for {
		fmt.Println("\nPlease select an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(tasks)
		case 2:
			listTasks(tasks)
		case 3:
			completeTask(tasks)
		case 4:
			saveTasks(tasks)
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func loadTasks() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}
	}
	return strings.Split(string(data), "\n")
}

func saveTasks(tasks []string) {
	data := strings.Join(tasks, "\n")
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func addTask(tasks []string) {
	var task string
	fmt.Print("Enter task: ")
	fmt.Scanln(&task)
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Println("Task added successfully!")
}

func listTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func completeTask(tasks []string) {
	listTasks(tasks)
	var index int
	fmt.Print("Enter the index of the task to complete: ")
	fmt.Scanln(&index)
	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid index")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	saveTasks(tasks)
	fmt.Println("Task completed!")
}
