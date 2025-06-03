package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/to404hanga/TaskTracker/internal/domain"
	"github.com/to404hanga/TaskTracker/internal/model"
	"github.com/to404hanga/TaskTracker/internal/service"
)

func help() {
	fmt.Println("Usage: go run main.go <command> [args]")
	fmt.Println("    add <description>         Add a new task.")
	fmt.Println("    update <id> <description> Update a task.")
	fmt.Println("    delete <id>               Delete a task.")
	fmt.Println("    list                      List all tasks.")
	fmt.Println("    list todo                 List all todo tasks.")
	fmt.Println("    list in-progress          List all in-progress tasks.")
	fmt.Println("    list done                 List all done tasks.")
	fmt.Println("    help                      Show this help message.")
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	svc := service.NewTaskService()
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("[syntax error]: `task-cli add` must have a description")
			return
		}
		if err := svc.AddTask(os.Args[2]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", svc.AutoIncrement)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("[syntax error]: `task-cli update` must have an ID and a description")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("[syntax error]: ID must be an integer")
			return
		}
		if err = svc.UpdateTask(id, os.Args[3]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Task updated successfully (ID: %d)\n", id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("[syntax error]: `task-cli delete` must have an ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("[syntax error]: ID must be an integer")
		}
		if err = svc.RemoveTask(id); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Task deleted successfully (ID: %d)\n", id)
	case "list":
		var res []domain.Task
		var err error
		if len(os.Args) > 2 {
			status := os.Args[2]
			if status != "todo" && status != "in-progress" && status != "done" {
				fmt.Println("[syntax error]: status must be one of `todo`, `in-progress`, `done`")
				return
			}
			res, err = svc.GetTasks(model.FromString(status))
		} else {
			res, err = svc.GetAllTasks()
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, task := range res {
			fmt.Printf("{ id: %d, description: %s status: %s, createdAt: %s, updatedAt: %s }\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("[syntax error]: `task-cli mark-in-progress` must have an ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("[syntax error]: ID must be an integer")
			return
		}
		if err = svc.MarkInProgress(id); err != nil {
			fmt.Println(err)
			return
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("[syntax error]: `task-cli mark-done` must have an ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("[syntax error]: ID must be an integer")
			return
		}
		if err = svc.MarkDone(id); err != nil {
			fmt.Println(err)
			return
		}
	default:
		help()
	}
}
