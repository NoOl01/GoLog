## GoLog
"GoLog" is a simple logging library based on the principle: one log — one file. With GoLog, you can customize certain features to suit your needs.

# Install
`go get github.com/NoOl01/GoLog`

# Features
- Log entries to both separate files and the console.
- Different types of logs: Error, Warning, Info.
- Files are created in the format: "error type — log date", which ensures easy searching by date.
- Option to disable console output.
- Ability to choose any name for the log folder.
- Option to change the file extension to something more convenient for you (the logging format remains unchanged and is recorded as a string, like in a regular text file).
- Cleanup of old logs with the option to set a custom date after which the log will be considered outdated.
- Timer setup for automatic log cleanup with a configurable interval.

# Known Issues
- In case of a large number of requests, the log may be overwritten instead of creating a new one.

# Example
```
package main

import (
	"errors"
	"github.com/NoOl01/GoLog/goLog"
	"github.com/NoOl01/GoLog/options"
)

func main() {
	//GoLog options
	options.SetLogDir("Your directory name")
	options.ToggleConsoleLog(false)            //Disables output to the console. by default: true
	options.SetFileExtension("Your extension") //By default: txt.
	options.Cleanup(30)                        //This file will be deleted after a set period of time.
	options.SetCleanupTime(1)                  // Runs a "Cleanup" every night after a set number of days. DOES NOT WORK WITHOUT CLEANUP

	//GoLog log functions
	goLog.ErrLog(errors.New("error"))
	goLog.WarnLog("warning")
	goLog.InfoLog("information")
}
```
