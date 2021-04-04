/*/
This file contains functions for printing in color and formatting
/*/
package go_practice

// import the libraries we need
import (
	// basic ANSI Escape sequences
	"github.com/fatih/color"
)

// Colorized error printing to see what we are doing
func error_printer(error_object error, message string) {

	//error_as_string, err := fmt.Errorf(error_object.Error())
	color.Red(error_object.Error(), message)
}

// debugging feedback function
// prints colored text for easy visual identification of data
func debug_print(pcolor string, message string) {
	switch pcolor {
	case pcolor == "red":
		color.Red(message)
	case pcolor == "blue":
		color.Blue(message)
	case pcolor == "green":
		color.Green(message)

	}
}
