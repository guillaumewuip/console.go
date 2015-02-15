package console

//Logger colors
//
//Uses colors from the https://github.com/mgutz/ansi package
//
//Default values are :
//	ColorsOptions{
// 	   "Log":     "green:black",
// 	   "Info":    "cyan:black",
// 	   "Error":   "red:black",
// 	   "Warning": "yellow:black",
//	}
//
type ColorsOptions map[string]string

var defaultColors = ColorsOptions{
	"Log":     "green:black",
	"Info":    "cyan:black",
	"Error":   "red:black",
	"Warning": "yellow:black",
}
