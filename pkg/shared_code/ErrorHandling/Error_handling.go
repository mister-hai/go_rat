/*/
This file contains functions for printing in color and formatting
/*/
package shared_code

// import the libraries we need
import (
	/*/  IMPORTING MODULES YOU FIND ONLINE
	in the terminal in VSCODE, while in the package root directory,
	append the following imports, as is, to the command "go get"

	Example:

	go get github.com/hashicorp/mdns

	And it will install the modules to the
	GOMODCACHE directory
	/*/
	"log"

	"github.com/fatih/color"
)

// this function is for backend usage, use the function
//rat_error()
// to replace the log and print for errors
// Colorized error printing to see what we are doing
// maybe have the error printer also take a struct, with the assignment
// being attempted, have the error assigned to it?
func Error_printer(error_object error, message string) {

	//error_as_string, err := fmt.Errorf(error_object.Error())
	color.Red(error_object.Error(), message)
}

// if "message" is "", it will simply log the error
// and respond as if it were the "log" function
func RatLogError(error_object error, message string) error {
	if message != "" {
		log.Printf(color.YellowString(error_object.Error()))
	}
	return error_object
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
