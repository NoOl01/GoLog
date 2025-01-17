package GoloG

import (
	"fmt"
	"github.com/NoOl01/GoLog/internal"
	"github.com/NoOl01/GoLog/options"

	"os"
	"time"
)

func writeLog(logType, logMessage string) {
	_, err := os.Stat(internal.DirectoryName)
	if err != nil {
		options.SetLogDir(internal.DirectoryName)
		fmt.Printf("Error: %s\n", err)
	}

	name := logType + "-" + time.Now().Format("02-01-2006_15.04")
	file, err := os.Create("./" + internal.DirectoryName + "/" + name + "." + string(internal.FileExtension))

	if err != nil {
		fmt.Printf("Error creating file: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v", err)
		}
	}(file)

	switch logType {
	case "err_log":
		{
			_, err = file.WriteString("[ERROR]: " + logMessage)
			if err != nil {
				return
			}
			if internal.Console {
				fmt.Printf("\033[31m[ERROR]\033[0m: %s\n", logMessage)
			}
		}
	case "warn_log":
		{
			_, err = file.WriteString("[WARNING]: " + logMessage)
			if err != nil {
				return
			}
			if internal.Console {
				fmt.Printf("\033[33m[WARNING]\033[0m: %s\n", logMessage)
			}
		}
	case "info_log":
		{
			_, err = file.WriteString("[INFO]: " + logMessage)
			if err != nil {
				return
			}
			if internal.Console {
				fmt.Printf("\033[37m[INFO]\033[0m: %s\n", logMessage)
			}
		}
	}
}

// ErrLog — outputs error logs to the console (if not disabled) and saves the logs to a new file with a date in the format "DD-MM-YYYY_hh.mm" (support for custom date format is under development).
func ErrLog(err error) {
	writeLog("err_log", err.Error())
}

// WarnLog — outputs warnings logs to the console (if not disabled) and saves the logs to a new file with a date in the format "DD-MM-YYYY_hh.mm" (support for custom date format is under development).
func WarnLog(warn string) {
	writeLog("warn_log", warn)
}

// InfoLog — outputs information logs to the console (if not disabled) and saves the logs to a new file with a date in the format "DD-MM-YYYY_hh.mm" (support for custom date format is under development).
func InfoLog(info string) {
	writeLog("info_log", info)
}
