package Cmd

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
	"time"
)

type Task struct {
	ID          int
	Text        string
	Status      string
	LastUpdated string
	TimeCreated string
}

func Add(path string, taskText string) {

	newID := getSmallestAvailableID(path)
	newTask := createTask(taskText, newID)

	fileWrite, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer fileWrite.Close()
	jsonData, err := json.Marshal(newTask)
	if err != nil {
		return
	}

	fileWrite.Write(jsonData)
	fileWrite.WriteString("\n")
}

func createTask(userInput string, newID int) Task {
	currentTime := time.Now().Format("02.01.2006 15:04")

	return Task{
		ID:          newID,
		Text:        userInput,
		Status:      "todo",
		TimeCreated: currentTime,
		LastUpdated: currentTime,
	}
}

func getSmallestAvailableID(path string) int {
	file, err := os.Open(path)
	if err != nil {
		return 1
	}
	defer file.Close()

	var ids []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t Task
		if err := json.Unmarshal(scanner.Bytes(), &t); err == nil {
			ids = append(ids, t.ID)
		}
	}

	sort.Ints(ids)

	smallestID := 1
	for _, id := range ids {
		if id == smallestID {
			smallestID++
		} else if id > smallestID {
			break
		}
	}
	return smallestID
}
