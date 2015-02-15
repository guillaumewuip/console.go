package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"path"
	"runtime"
	"strconv"
	"time"
)

//Logger options
type Options struct {
	Colors            ColorsOptions
	ContextMediumSize int
	SpaceSize         int
	DefaultTags       []string
}

type Show struct {
	Tags     bool
	Location bool
	Time     bool
}

//A log location
type Location struct {
	Filename string
	Line     int
}

//Logger struct stores log params
//up to the print
type Logger struct {
	Level          string
	Tags           []string
	Timestamp      int64
	ContextMessage string
	Message        string
	Hooks          internalHooks
	Show
	Location
	Options
}

//Reset the logger after print
func (logger *Logger) reset() {

	logger.Tags = []string{}
	logger.Show = Show{}

}

//Build the context string
//ie. tags, time, location
func (logger *Logger) buildContext() {

	var context string

	//Build location
	_, logger.Location.Filename, logger.Location.Line, _ = runtime.Caller(3)
	logger.Location.Filename = path.Base(logger.Location.Filename)

	//Build time
	logger.Timestamp = time.Now().Unix()

	//Build context message

	if logger.Show.Tags {
		for _, tag := range logger.Tags {
			context += "[" + tag + "]"
		}
	}

	if logger.Show.Location {
		context += " [" + logger.Location.Filename + ":" + strconv.Itoa(logger.Location.Line) + "] "
	}

	if logger.Show.Time {
		context += time.Now().Format(time.RFC3339) + " "
	}

	logger.ContextMessage = context
}

//Print the log in console
//And fire hooks
func (logger *Logger) printLog(color string, msg string, args ...interface{}) {

	if len(args) > 0 {
		logger.Message = fmt.Sprintf(msg, args...)
	} else {
		logger.Message = msg
	}

	//Build context string
	logger.buildContext()

	//Fire hook
	logger.Hooks.Fire(*logger)

	//Print message to console
	fmt.Println(ansi.Color(
		logger.ContextMessage+logger.Message,
		color,
	))

	logger.reset()
}

//Add tags to the log
func (logger *Logger) Tag(args ...string) *Logger {

	logger.Show.Tags = true
	logger.Tags = append(logger.Tags, args...)

	return logger
}

//Add file and line information to the log
func (logger *Logger) File() *Logger {

	logger.Show.Location = true

	return logger
}

//Add time information to the log
func (logger *Logger) Time() *Logger {
	logger.Show.Time = true

	return logger
}

//Log functions

//Log level
//	console.Log("Hello World")
func (logger *Logger) Log(msg string, args ...interface{}) {
	logger.Level = "log"
	color := logger.Options.Colors["Log"]
	logger.printLog(color, msg, args...)
}

//Info level
//	console.Info("Hello World")
func (logger *Logger) Info(msg string, args ...interface{}) {
	logger.Level = "info"
	color := logger.Options.Colors["Info"]
	logger.printLog(color, msg, args...)
}

//Error level
//	console.Error("Hello World")
func (logger *Logger) Error(msg string, args ...interface{}) {
	logger.Level = "error"
	color := logger.Options.Colors["Error"]
	logger.printLog(color, msg, args...)
}

//Warn level
//	console.Warn("Hello World")
func (logger *Logger) Warning(msg string, args ...interface{}) {
	logger.Level = "warning"
	color := logger.Options.Colors["Warning"]
	logger.printLog(color, msg, args...)
}
