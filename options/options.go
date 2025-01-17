package options

import (
	"fmt"
	"github.com/NoOl01/GoLog/internal"
	"os"
)

// SetLogDir — sets the name of the folder for storing logs.
// `options.SetLogDir(<Your directory name>)`
func SetLogDir(dirName string) {
	_, err := os.Stat(dirName)
	if err == nil {
		fmt.Printf("[GoLog INFO] Directory %s already exist", dirName)
		internal.DirectoryName = dirName
		return
	}
	if dirName == "" {
		dirName = "logger"
	}
	err = os.Mkdir("./"+dirName, 07)
	if err != nil {
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
