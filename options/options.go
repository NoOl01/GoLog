package options

import (
	"errors"
	"fmt"
	"os"
)

var DirectoryName = "logger"

//TODO
// var TimeZone int8 = 0

var Console = true
var FileExtension Extension = TxtExtension

type Extension string

const (
	TxtExtension Extension = "txt"
	LogExtension Extension = "log"
)

// SetLogDir — sets the name of the folder for storing logs.
// `options.SetLogDir(<Your directory name>)`
func SetLogDir(dirName string) {
	_, err := os.Stat(dirName)
	if err == nil {
		fmt.Printf("[LoGo INFO] Directory %s already exist", dirName)
		return
	}
	if dirName == "" {
		dirName = "logger"
	}
	err = os.Mkdir("./"+dirName, 0777)
	if err != nil {
		return
	}
	DirectoryName = dirName
}

//TODO
//func SetTimeZone(timeZone int8) error {
//	if timeZone < -12 || timeZone > 12 {
//		return fmt.Errorf("invalid TimeZone value: %d; must be in range [-12, 12]", timeZone)
//	}
//	TimeZone = timeZone
//	return nil
//}

// ToggleConsoleLog — toggle logging output to the console.
// by default - true.
// `options.ToggleConsoleLog(false)`
func ToggleConsoleLog(consoleOut bool) {
	Console = consoleOut
}

// SetFileExtension — Allows changing the file extension used to save logs. (Currently, only two extensions are available: ".txt" and ".log")
func SetFileExtension(extension Extension) error {
	switch extension {
	case TxtExtension, LogExtension:
		FileExtension = extension
		return nil
	default:
		return errors.New("invalid file extension: only 'txt' and 'log' are allowed")
	}
}
