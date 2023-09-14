package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// CHANGE FILE PATH ACCORDINGLY
	folderPath := "enter file path here"

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		oldName := file.Name()
		newName := strings.ReplaceAll(oldName, " ", "") // Remove spaces from the file name

		if oldName != newName {
			oldPath := filepath.Join(folderPath, oldName)
			newPath := filepath.Join(folderPath, newName)

			if err := os.Rename(oldPath, newPath); err != nil {
				fmt.Printf("Error renaming %s to %s: %v\n", oldPath, newPath, err)
			} else {
				fmt.Printf("Renamed %s to %s\n", oldPath, newPath)
			}
		}
	}
}
