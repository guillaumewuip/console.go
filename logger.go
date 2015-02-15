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
type options struct {
	Colors            ColorsOptions
	ContextMediumSize int
	SpaceSize         int
	DefaultTags       []string
}

type show struct {
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
	hooks          internalHooks
	Location
	show
	options
}

//Reset the logger after print
func (logger *Logger) reset() {

	logger.Tags = []string{}
	logger.show = show{}

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

	if logger.show.Tags {
		for _, tag := range logger.Tags {
			context += "[" + tag + "]"
		}
	}

	if logger.show.Location {
		context += " [" + logger.Location.Filename + ":" + strconv.Itoa(logger.Location.Line) + "] "
	}

	if logger.show.Time {
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
	logger.hooks.Fire(*logger)

	//Print message to console
	fmt.Println(ansi.Color(
		logger.ContextMessage+logger.Message,
		color,
	))

	logger.reset()
}

//Add tags to the log
func (logger *Logger) Tag(args ...string) *Logger {

	logger.show.Tags = true
	logger.Tags = append(logger.Tags, args...)

	return logger
}

//Add file and line information to the log
func (logger *Logger) File() *Logger {

	logger.show.Location = true

	return logger
}

//Add time information to the log
func (logger *Logger) Time() *Logger {
	logger.show.Time = true

	return logger
}

//Log functions

//Log level
//	console.Log("Hello World")
func (logger *Logger) Log(msg string, args ...interface{}) {
	logger.Level = "log"
	color := logger.options.Colors["Log"]
	logger.printLog(color, msg, args...)
}

//Info level
//	console.Info("Hello World")
func (logger *Logger) Info(msg string, args ...interface{}) {
	logger.Level = "info"
	color := logger.options.Colors["Info"]
	logger.printLog(color, msg, args...)
}

//Error level
//	console.Error("Hello World")
func (logger *Logger) Error(msg string, args ...interface{}) {
	logger.Level = "error"
	color := logger.options.Colors["Error"]
	logger.printLog(color, msg, args...)
}

//Warn level
//	console.Warn("Hello World")
func (logger *Logger) Warning(msg string, args ...interface{}) {
	logger.Level = "warning"
	color := logger.options.Colors["Warning"]
	logger.printLog(color, msg, args...)
}
