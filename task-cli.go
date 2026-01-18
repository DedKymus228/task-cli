package main

import (
	"Task_Tracker/Cmd"
	"Task_Tracker/Data"
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := Data.CheckFile()

	if len(os.Args) < 2 {
		fmt.Println("Enter help")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: add <'task'>")
			return
		}
		taskText := os.Args[2]
		Cmd.Add(path, taskText)

	case "mark-in-progress", "mark-done":
		command := os.Args[1]
		if len(os.Args) < 3 {
			fmt.Printf("Error: %s <ID>", command)
		}
		targetId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:not number")
			return
		}
		Cmd.Mark(path, command, targetId)

	case "list":
		status := os.Args[2]
		Cmd.List(path, status)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: update <ID> <'task'>")
			return
		}
		targetId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:not number")
			return
		}
		newText := os.Args[3]
		Cmd.Update(path, targetId, newText)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: delete <ID>")
			return
		}
		targetId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:not number")
			return
		}
		Cmd.Delete(path, targetId)

	case "help":
		fmt.Println("\n=== Task Tracker CLI Help ===")
		fmt.Println("\nUsage:")
		fmt.Println("  go run task-cli.go <command> [arguments]")

		fmt.Println("\nAvailable Commands:")

		fmt.Println("  add \"task\"\t\tAdd a new task")
		fmt.Println("  list\t\t\tList all tasks")
		fmt.Println("  list done\t\tList only completed tasks")
		fmt.Println("  list todo\t\tList tasks that are not done")
		fmt.Println("  list in-progress\tList tasks currently being worked on")
		fmt.Println("  update <id> \"task\"\tUpdate task description")
		fmt.Println("  delete <id>\t\tRemove a task by its ID")
		fmt.Println("  mark-done <id>\tMark a specific task as completed")

		fmt.Println("\nExamples:")
		fmt.Println("  go run task-cli.go add \"Buy some milk\"")
		fmt.Println("  go run task-cli.go list done")
		fmt.Println("  go run task-cli.go update 1 \"Buy milk and bread\"")

	default:
		fmt.Println("no com", command)
	}
}
