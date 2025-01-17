package options

import (
	"fmt"
	"github.com/NoOl01/GoLog/internal"
	"log"
	"os"
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
