package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var tasks []task
var loadedTask []task

func main() {
	loadTask()
	if len(os.Args) < 2 {
		fmt.Println("please enter a command")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("There must be a command")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		add(description)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("please enter a command")
			return
		}
		idstr := os.Args[2]
		deleteTask(idstr)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("add a new description")
			return
		}
		idstr := os.Args[2]
		newdesc := strings.Join(os.Args[3:], " ")
		update(idstr, newdesc)
	case "mark":
		if len(os.Args) < 4 {
			fmt.Println("not enough commands")
			return
		}
		idstr := os.Args[2]
		newStatus := strings.Join(os.Args[3:], " ")
		mark(idstr, newStatus)
	case "list":
		listAll()
	case "list-not-done":
		listNotDone()
	case "list-in-progress":
		listTaskInprogress()
	default:
		fmt.Println("Unknown command. Use add, delete, update, mark, list, ...")

	}
}

func add(description string) {
	newTask := task{
		ID:          geNnextID(),
		Description: description,
		Status:      "todo",
	}

	tasks = append(tasks, newTask)
	saveTask()
	fmt.Println("task successfully added")
	fmt.Printf("Added task: ID:%d, Description:%s, Status:%s\n", newTask.ID, newTask.Description, newTask.Status)
}
func deleteTask(idstr string) {
	var update []task
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("There was an error trying to convert")
	}
	found := false
	for _, r := range tasks {
		if r.ID != id {
			update = append(update, r)
		} else {
			found = true
		}

	}
	if !found {
		fmt.Println("task does not have an id")
		return
	}
	tasks = update
	saveTask()
	fmt.Println("tasks successfully deleted")
	listAll()
}
func update(idstr, newDesc string) {
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(" There was an error trying to convert")
		return
	}
	found := false
	for i, r := range tasks {
		if r.ID == id {
			tasks[i].Description = newDesc
			found = true
		}
	}
	if !found {
		fmt.Println("ID was not found")
		return
	}
	fmt.Println("Task successfully updated")
	saveTask()
	listAll()
}
func mark(idstr, newStatus string) {
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Error in converting to integer")
		return
	}
	found := false
	validStatuses := map[string]bool{"todo": true, "in-progress": true, "done": true}
	if !validStatuses[newStatus] {
		fmt.Println("Invalid status. Use: todo, in-progress, done")
		return
	}
	for i, r := range tasks {
		if r.ID == id {
			tasks[i].Status = newStatus
			found = true
			break
		}

	}
	if !found {
		fmt.Println("ID not found")
		return
	}
	saveTask()
	fmt.Println("status changed")
}
func listAll() {
	for _, r := range tasks {
		fmt.Printf("ID:%d, Description:%s,Status:%s\n", r.ID, r.Description, r.Status)
	}
}
func listNotDone() {
	var notDone []task

	for _, r := range tasks {
		if r.Status == "todo" || r.Status == "in-progress" {
			notDone = append(notDone, r)

		}
	}
	if len(notDone) == 0 {
		fmt.Println("no task found that are not done")
		return
	}
	for _, t := range notDone {
		fmt.Printf("ID:%d, Description:%s, Status:%s\n", t.ID, t.Description, t.Status)
	}

}
func listTaskInprogress() {
	var taskInprogress []task

	for _, r := range tasks {
		if r.Status == "in-progress" {
			taskInprogress = append(taskInprogress, r)

		}
	}
	if len(taskInprogress) == 0 {
		fmt.Println("no tasks found that are in-progress")
		return
	}
	for _, t := range taskInprogress {
		fmt.Printf("ID:%d, Description:%s, Status:%s\n", t.ID, t.Description, t.Status)
	}

}
func saveTask() {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error when converting to json")
		return
	}
	os.WriteFile("tasks.json", data, 0644)
	fmt.Println("file has been saved")
}
func loadTask() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("error reading file")
	}
	json.Unmarshal(data, &loadedTask)
	tasks = loadedTask
}
func geNnextID() int {
	maxId := 0
	for _, r := range tasks {
		if r.ID > maxId {
			maxId = r.ID
		}
	}
	return maxId + 1
}
