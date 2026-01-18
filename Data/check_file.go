package Data

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckFile() string {
	folder := "Data"
	fileName := "Tasks.json"

	path := filepath.Join(folder, fileName)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.MkdirAll(folder, 0755)
		file, _ := os.Create(path)
		defer file.Close()
	} else if err != nil {
		fmt.Println("Error:", err)
	}
	return path
}
