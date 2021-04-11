/*/
This file contains functions for printing in color and formatting
/*/
package ErrorHandling

// import the libraries we need
import (
	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
)

// this function is for backend usage, use the function
//RatLogError()
// to replace the log and print for errors
func ErrorPrinter(derp error, message string) error {

	//error_as_string, err := fmt.Errorf(error_object.Error())
	color.Red(derp.Error(), message)
	return derp
}

// use this instead of the regular logging functions
// returns the errors but adds them to a log while printing a
// message to the screen for your viewing pleasure
// responds to log level changes, will affect other logging activities
func RatLogError(derp error, message string, logfile string) error {
	log.SetFormatter(&log.JSONFormatter{})
	// print to terminal for debug purposes
	// we can do either:
	//logs, derp := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return derp
}

func ShowLogs(LinesToPrint int, loglevel string) {
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
