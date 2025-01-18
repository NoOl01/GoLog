package options

import (
	"fmt"
	"github.com/NoOl01/GoLog/internal"
	"log"
	"os"
	"sync"
	"time"
)

// SetLogDir — sets the name of the folder for storing logs.
// `options.SetLogDir(<Your directory name>)`
func SetLogDir(dirName string) {
	dir, _ := os.Stat(dirName)
	if dir != nil && dir.IsDir() {
		fmt.Printf("\u001B[37m[GoLog INFO]\u001B[0m Directory %s already exist\n", dirName)
	}
	if dirName == "" {
		dirName = "logger"
	}
	err := os.Mkdir("./"+dirName, 0755)
	if err != nil {
		log.Panicf("\u001B[31m[GoLog ERROR]\u001B[0m: %s", err)
		return
	}
	internal.DirectoryName = dirName
}

//TODO:
// func SetTimeZone(timeZone int8) error {
//	if timeZone < -12 || timeZone > 12 {
//		return fmt.Errorf("invalid TimeZone value: %d; must be in range [-12, 12]", timeZone)
//	}
//	TimeZone = timeZone
//	return nil
// }

// ToggleConsoleLog — toggle logging output to the console.
// by default - true.
// `options.ToggleConsoleLog(false)`
func ToggleConsoleLog(consoleOut bool) {
	internal.Console = consoleOut
}

// SetFileExtension — Allows changing the file extension used to save logs.
func SetFileExtension(extension string) {
	if extension == "" {
		internal.FileExtension = "txt"
	}
	internal.FileExtension = extension
}

// Cleanup is a function that automatically deletes logs after the specified period has passed.
//
// The "days" parameter defines the number of days after which the logs will be deleted. The formula for calculating it is as follows: "currentDay - <your day value>".
//
// For example, if the current day is 2025-01-18 and the deadline is 2025-01-17, then files with a date that is less than or equal to the deadline will be deleted.
func Cleanup(days int16) {
	if days == 0 {
		fmt.Printf("\u001B[37m[GoLog INFO]\u001B[0m: the set day cannot be equal to 0\n")
		return
	}
	internal.CleanerDeadline = days
	deadLine := time.Now().Add(time.Duration(-days) * 24 * time.Hour).Format("2006-01-02")

	wg := sync.WaitGroup{}
	wg.Add(3)

	filesChan := make(chan os.DirEntry)
	outdated := make(chan os.DirEntry)

	go func() {
		defer wg.Done()
		internal.SearchFiles(filesChan)
		close(filesChan)
	}()

	go func() {
		defer wg.Done()
		internal.OutdatedFiles(filesChan, outdated, deadLine)
		close(outdated)
	}()

	go func() {
		defer wg.Done()
		internal.DeleteOutdatedFiles(outdated)
	}()

	wg.Wait()
}

// SetCleanupTime is a function that sets a timer to call "Cleanup" every night at 00:00. The "days" parameter defines the interval after which the function will be triggered.
//
// For example, if the "days" value is set to 2, the function will execute every 2 days.
//
// !IMPORTANT! For the timer to work correctly, you must first use the `Cleanup` function.
func SetCleanupTime(days int) {
	timer(days)
}

func timer(day int) {
	currentTime := time.Now()
	nextNight := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+day, 0, 0, 0, 0, currentTime.Location())

	duration := nextNight.Sub(currentTime)

	time.AfterFunc(duration, func() {
		Cleanup(internal.CleanerDeadline)

		ticker := time.NewTicker(time.Duration(day) * 24 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			Cleanup(internal.CleanerDeadline)
		}
	})
}
