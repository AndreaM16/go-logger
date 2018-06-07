// Package log handles all the logs of the application.
// Only use its public Print functions to print the different logs.
// Log Types:
//     - Info: Trace logs
//     - Warning: Something is wrong but not blocking
//     - Error: Something is wrong and blocking
// Output Examples:
//     - INFO: 2013/11/05 18:11:01 main.go:44: Special Information
//     - WARNING: 2013/11/05 18:11:01 main.go:45: There is something you need to know about
//     - ERROR: 2013/11/05 18:11:01 main.go:46: Something has failed
// Functions:
//    - Prints: Normal prints accepting a string or an error as input
//    - PrintFs: Prints accepting a string or an error + variadic arguments as input. They format the arguments
//      in the passed string
// Function Calls Examples:
//    - log.Info.Print("Just one info")
//    - log.Info.Printf("Info number %v", 1)
package log

import (
    "io"
    "log"
    "os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// init initializes the standard output
func init() {
	Init(os.Stdout, os.Stdout, os.Stderr)
}

// Init initializes the standard output for each kind of log
func Init(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
