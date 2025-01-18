package internal

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func SearchFiles(filesChan chan os.DirEntry) {
	dir, _ := os.Stat(DirectoryName)
	if dir == nil || !dir.IsDir() {
		fmt.Printf("\u001B[31m[GoLog ERROR]\u001B[0m: Directory '%s' does not exist.\n", DirectoryName)
		return
	}
	files, err := os.ReadDir(DirectoryName)
	if err != nil {
		fmt.Printf("\u001B[31m[GoLog ERROR]\u001B[0m: %s\n", DirectoryName)
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			filesChan <- file
		}
	}
}

func OutdatedFiles(filesChan chan os.DirEntry, outdated chan os.DirEntry, deadLine string) {
	layout := "2006-01-02"

	parsedDl, err := time.Parse(layout, deadLine)
	if err != nil {
		fmt.Printf("\u001B[31m[GoLog ERROR]\u001B[0m: %s\n", err)
		return
	}

	for file := range filesChan {
		splitFile := strings.Split(file.Name(), ".")
		parsedTime, err := time.Parse(layout, splitFile[1])
		if err != nil {
			fmt.Printf("\u001B[31m[GoLog ERROR]\u001B[0m: %s\n", err)
			break
		}
		if parsedTime.Before(parsedDl) || parsedTime.Equal(parsedDl) {
			outdated <- file
		}
	}
}

func DeleteOutdatedFiles(outdated chan os.DirEntry) {
	for file := range outdated {
		err := os.Remove("./" + DirectoryName + "/" + file.Name())
		if err != nil {
			fmt.Printf("\u001B[31m[GoLog ERROR]\u001B[0m: %s\n", err)
			break
		}
	}
}
