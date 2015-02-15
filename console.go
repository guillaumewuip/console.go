//
//Console is a lightweight logging tool inspired by the NodeJS console object
// and http://github.com/bluejamesbond/scribe.js project.
//
//	console.Info("Some info")
//	console.Tag("Hello").Time().File().Log("Hello %s", "World")
//
//Console provided Info, Log, Warning and Error method and Tag, Time, File optional information
//
//With use of ansi colors, it output something like below in terminal :
//	[Tag1][Tag2] [main.go:48] 2015-02-15T09:22:06+01:00 Hello World
//
//Author: http://github.com/guillaumewuip
package console

//Create a new console
//	myConsole = console.NewConsole(console.ColorsOptions{})
//Pass a ColorsOptions map (could be empty to use default values).
//
//Colors from https://github.com/mgutz/ansi.
//See examples below :
func NewConsole(colors ColorsOptions) *Logger {

	logger := Logger{}
	logger.options.Colors = colors

	for key, color := range defaultColors {
		if colors[key] == "" {
			colors[key] = color
		}
	}

	return &logger
}
