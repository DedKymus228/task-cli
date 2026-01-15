package Cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func List(path string, status string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:No tasks")
	}

	scanner := bufio.NewScanner(file)
	if status != "list" {
		for scanner.Scan() {
			var t Task
			json.Unmarshal(scanner.Bytes(), &t)
			if status == t.Status {
				fmt.Println(t.Text)
			}
		}
	} else {
		for scanner.Scan() {
			var t Task
			json.Unmarshal(scanner.Bytes(), &t)
			fmt.Printf(t.Text)
		}

	}
	defer file.Close()
}
