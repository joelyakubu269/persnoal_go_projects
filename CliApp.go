package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int
	Description string
	Status      string
	//createdAt    time.time
	//UpdatedAt    time.time

}

var nextID = 1
var Tasks []Task

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please enter a command")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("please enter a task")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		add(description)
	case "list":
		listTasks()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("please provide a task ID")
			return
		}
		deleteTask(os.Args[2])
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("plrease provide task ID")
			return
		}
		completeTask(os.Args[2])
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("please provide ID and new description")
			return
		}
		idstr := os.Args[2]
		newDesc := strings.Join(os.Args[3:], " ")
		updateTask(idstr, newDesc)
	default:
		fmt.Println("please do something")
		return
	}

}
func add(description string) {
	newTask := Task{
		ID:          nextID,
		Description: description,
		Status:      "todo",
	}
	Tasks = append(Tasks, newTask)
	nextID++
}
func listTasks() {
	if len(Tasks) < 1 {
		fmt.Println("There is nothing to add")
		return
	}
	for _, r := range Tasks {
		fmt.Printf("ID:%d ,%s,%s\n", r.ID, r.Description, r.Status)
	}
}
func deleteTask(idstr string) {
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	var updated []Task
	found := false
	for _, d := range Tasks {
		if d.ID != id {
			updated = append(updated, d)
		} else {
			found = true
		}
	}
	if !found {
		fmt.Println("ID is not found")
		return
	}
	Tasks = updated
	fmt.Println("Task deleted")
}
func completeTask(idstr string) {
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	for i, t := range Tasks {
		if err != nil {
			if t.ID == id {
				Tasks[i].Status = "done"
				fmt.Println("Task marked as complete")
				return
			}
		}
		fmt.Println("Task not found")
	}
}
func updateTask(idstr, newDesc string) {
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Invaalid ID")
		return
	}
	for i, t := range Tasks {
		if t.ID == id {
			Tasks[i].Description = newDesc
			fmt.Println("Task updated")
			return
		}
	}
	fmt.Println("task not found")
}
