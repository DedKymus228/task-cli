package Cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func Mark(path string, command string, targetId int) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var allTasks []Task

	for scanner.Scan() {
		var t Task
		json.Unmarshal(scanner.Bytes(), &t)

		if t.ID == targetId {
			if command == "mark-in-progress" {
				t.Status = "in-progress"
			} else if command == "mark-done" {
				t.Status = "done"
			} else {
				fmt.Println("error")
			}
			t.LastUpdated = time.Now().Format("02.01.2006 15:04")
		}
		allTasks = append(allTasks, t)
	}
	file.Close()

	newFile, _ := os.Create(path)
	defer newFile.Close()

	for _, t := range allTasks {
		jsonData, _ := json.Marshal(t)
		newFile.Write(jsonData)
		newFile.WriteString("\n")
	}
}
