package goLog

import (
	"fmt"
	"github.com/NoOl01/GoLog/internal"
	"github.com/NoOl01/GoLog/options"
	"os"
	"time"
)

func writeLog(logType, logMessage string) {
	dir, _ := os.Stat(internal.DirectoryName)
	if dir == nil || !dir.IsDir() {
		options.SetLogDir(internal.DirectoryName)
		fmt.Printf("\u001B[37m[GoLog INFO]\u001B[0m: %s not exist. Creating directory\n", internal.DirectoryName)
	}

	name := logType + "-" + time.Now().Format("02-01-2006_15.04")
	file, err := os.Create("./" + internal.DirectoryName + "/" + name + "." + string(internal.FileExtension))

	if err != nil {
		fmt.Printf("\u001B[31m[GoLog ERROR]\u001B[0m: %v", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("\u001B[31m[GoLog ERROR]\u001B[0m: %v", err)
			return
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
				fmt.Printf("\u001B[31m[ERROR]\u001B[0m: %s\n", logMessage)
			}
		}
	case "warn_log":
		{
			_, err = file.WriteString("[WARNING]: " + logMessage)
			if err != nil {
				return
			}
			if internal.Console {
				fmt.Printf("\u001B[31m[ERROR]\u001B[0m: %s\n", logMessage)
			}
		}
	case "info_log":
		{
			_, err = file.WriteString("[INFO]: " + logMessage)
			if err != nil {
				return
			}
			if internal.Console {
				fmt.Printf("\u001B[31m[ERROR]\u001B[0m: %s\n", logMessage)
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
