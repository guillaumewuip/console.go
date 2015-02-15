package console_test

//Create a new console with default colors
func ExampleNewConsole() {
	myConsole := console.NewConsole(console.ColorsOptions{})
}

//Create a new console with custom colors
func ExampleNewConsole_withCustomColors() {
	colors := console.ColorsOptions{
		"Log": "green+i:black",
	}

	myConsole := console.NewConsole(colors)
}

//Here's an example showing how to received all log.
func ExampleHookFunc() {

	testHook := func(logger console.Logger) error {

		fmt.Printf(
			"Test hook fired. Message=%s, Level=%s, File=%s \n",
			logger.Message,
			logger.Level,
			logger.Location.Filename,
		)

		return nil
	}

	myConsole.AddHook(testHook)

}

//Here's an example showing how to received all log.
func ExampleAddHook() {

	testHook := func(logger console.Logger) error {

		fmt.Printf(
			"Test hook fired. Message=%s, Level=%s, File=%s \n",
			logger.Message,
			logger.Level,
			logger.Location.Filename,
		)

		return nil
	}

	myConsole.AddHook(testHook)

}
