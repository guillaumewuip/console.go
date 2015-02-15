package main

import (
	K "github.com/guillaumewuip/console.go"
)

func main() {

	colors := K.ColorsOptions{
		"Log": "green+i:black",
	}
	console := K.NewConsole(colors)

	//Add hook
	testHook := func(logger K.Logger) error {

		fmt.Printf(
			"Test hook fired. Message=%s, Level=%s, File=%s \n",
			logger.Message,
			logger.Level,
			logger.Location.Filename,
		)

		return nil
	}

	console.AddHook(testHook)

	world := "World"

	console.Tag("First", "Test").Time().File().Log("Hello %s", world)
	console.Info("Info")
	console.Error("Error")
	console.Warning("Warn")
}
