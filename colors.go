package console

type ColorsOptions map[string]string

var defaultColors = ColorsOptions{
	"Log":     "green:black",
	"Info":    "cyan:black",
	"Error":   "red:black",
	"Warning": "yellow:black",
}
