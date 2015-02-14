//Console is a lightweight logging tool
//	log.Tag("Hello").Time().File().Log("Hello World")
package console

//Create a new console
//	myConsole = console.NewConsole(console.ColorsOptions{})
//Pass a ColorsOptions map (could be empty to use default values)
//Colors from https://github.com/mgutz/ansi
func NewConsole(colors ColorsOptions) *Logger {

	logger := Logger{}
	logger.Options.Colors = colors

	for key, color := range defaultColors {
		if colors[key] == "" {
			colors[key] = color
		}
	}

	return &logger
}
