package console

//A Hook implements a function fired each time a log is printed.
//It receives a Logger.
type HookFunc func(Logger) error

//Internal type to store hooks in a Logger
type internalHooks []HookFunc

func (hooks internalHooks) Fire(logger Logger) {

	for _, hook := range hooks {
		hook(logger)
	}
}

//Method to add a hook to a console
func (logger *Logger) AddHook(hook HookFunc) {
	logger.hooks = append(logger.hooks, hook)
}
