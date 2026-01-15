package Cmd

import (
	"bufio"
	"encoding/json"
	"os"
)

func Delete(path string, targetId int) {
	file, _ := os.Open(path)
	var remainingTasks []Task
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var t Task
		err := json.Unmarshal(scanner.Bytes(), &t)
		if err != nil {
			continue
		}
		if targetId != t.ID {
			remainingTasks = append(remainingTasks, t)
		}
	}
	file.Close()
	newFile, err := os.Create(path)
	if err != nil {
		return
	}
	defer newFile.Close()
	for _, t := range remainingTasks {

		jsonData, _ := json.Marshal(t)
		newFile.Write(jsonData)
		newFile.WriteString("\n")
	}
}
