/*/
This file contains functions for printing in color and formatting
/*/
package ErrorHandling

// import the libraries we need
import (
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

// if "message" is "", it will simply log the error
// and respond as if it were the "log" function
func RatLogError(derp error, message string) error {
	if message != "" {
		log.Print(color.YellowString(derp.Error()))
	}
	return derp
}

func ShowLogs(LinesToPrint int) {
	log.SetFormatter(&log.JSONFormatter{})
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
