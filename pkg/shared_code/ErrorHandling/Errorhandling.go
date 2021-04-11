/*/
This file contains functions for printing in color and formatting
/*/
package ErrorHandling

// import the libraries we need
import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
)

// this function is for backend usage, use the function
//RatLogError()
// to replace the log and print for errors
func ErrorPrinter(derp error, message string) error {
	color.Red(derp.Error(), message)
	return derp
}

// use this to start the logger, cant keep it in globals.go
// returns 0 if failure to open logfile, returns 1 otherwise
// uses code from :
// https://esc.sh/blog/golang-logging-using-logrus/
func StartLogger(logfile string) (return_code int) {
	Logs, derp := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	LoggerInstance := log.New()
	Formatter := new(log.TextFormatter)
	Formatter.ForceColors = true
	Formatter.FullTimestamp = true
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	LoggerInstance.SetFormatter(Formatter)
	if derp != nil {
		// Cannot open log file. Logging to stderr
		ErrorPrinter(derp, "[-] ERROR: Failure To Open Logfile!")
		return 0
	} else {
		log.SetOutput(Logs)
	}
	return 1
}

// use this instead of the regular logging functions
// ONLY use for errors!
// returns the errors but adds them to a log while printing a
// message to the screen for your viewing pleasure
func RatLogError(derp error, message string) error {
	// output is being redirected to a file so we have to print as well
	log.Error(message)
	color.Red(message)
	return derp
}

// shows entries from the logfile, starting at the bottom
// limit by line number, loglevel, or time
func ShowLogs(LinesToPrint int, loglevel string, time string) {
	// set log thingie to use json so our json file
	// can be used
	log.SetFormatter(&log.JSONFormatter{})
	//switch loglevel{
	//	case "error":
	//		log.ErrorLevel

}

// debugging feedback function
// prints colored text for easy visual identification of data
// color_int (1,red)(2,green)(3,blue)(4,yellow)
func Debug_print(color_int int8, message string) {
	//if color_int
	switch color_int {
	//is 1
	case 1:
		color.Red(message)
		// and so on
	case 2:
		color.Blue(message)
	case 3:
		color.Green(message)
	case 4:
		color.Yellow(message)
	}
}
