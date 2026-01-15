package Cmd

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
)

func Update(path string, targetId int, newText string) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	var allTasks []Task

	for scanner.Scan() {
		var t Task
		json.Unmarshal(scanner.Bytes(), &t)

		if t.ID == targetId {
			t.Text = newText
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
