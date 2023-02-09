// Package slug is a simple structured logging utility.
package slug

import (
	"github.com/byReqz/slug/console"
	"os"
)

var DefaultLoggerSet = newDefaultSet() // logger set using default configuration

// newDefaultSet sets the default LoggerSet from console.NewDefaultConsoleLoggers
func newDefaultSet() *LoggerSet {
	ls := NewLoggerSet()
	for _, l := range console.NewDefaultConsoleLoggers() {
		ls.AddLogger(l)
	}
	return &ls
}

var (
	NoLevel      = -2 // log without level, level-less logs will always be printed regardless of logger level
	DebugLevel   = -1 // log at >= debug level
	InfoLevel    = 0  // log at >= info level
	WarningLevel = 1  // log at >= warning level
	ErrorLevel   = 2  // log at >= error level
	Disabled     = 3  // log at > error level (disable the logger)
)

// Logger defines a standardized logger interface.
type Logger interface {
	Level() int                // return the loggers level
	SWrite(data ...any) string // return a string formatted as the logger
	Write(data ...any)         // write data to the logger
	WriteE(data ...any) error  // write data to the logger and report about errors
}

// LoggerSet is a set of loggers.
type LoggerSet struct {
	Level   int      // the level which the set is operating on
	loggers []Logger // the loggers of the set
}

// NewLoggerSet returns a LoggerSet with the provided loggers.
func NewLoggerSet(loggers ...Logger) LoggerSet {
	return LoggerSet{loggers: loggers}
}

// AddLogger adds a new logger to the set.
func (ls *LoggerSet) AddLogger(l ...Logger) {
	ls.loggers = append(ls.loggers, l...)
}

// SprintAt returns the formatted outputs of the sets loggers at the given loglevel.
func (ls *LoggerSet) SprintAt(level int, v ...any) []string {
	var outs []string
	for _, l := range ls.loggers {
		if level == l.Level() {
			outs = append(outs, l.SWrite(v...))
		}
	}
	return outs
}

// PrintAt prints with all the sets loggers at the given loglevel.
func (ls *LoggerSet) PrintAt(level int, v ...any) {
	for _, l := range ls.loggers {
		if level == l.Level() {
			l.Write(v...)
		}
	}
}

// PrintlnAt prints with all the sets loggers at the given loglevel (and adds a newline).
func (ls *LoggerSet) PrintlnAt(level int, v ...any) {
	ls.PrintAt(level, append(v, "\n")...)
}

// Println prints a level-less log-entry to all loggers in the set (and adds a newline).
func (ls *LoggerSet) Println(v ...any) {
	ls.PrintlnAt(NoLevel, v...)
}

// Print prints a level-less log-entry to all loggers in the set.
func (ls *LoggerSet) Print(v ...any) {
	ls.PrintAt(NoLevel, v...)
}

// Sprint returns the level-less log-entries to all loggers in the set.
func (ls *LoggerSet) Sprint(v ...any) {
	ls.SprintAt(NoLevel, v...)
}

// Debug prints a debug-level log-entry to all loggers in the set.
func (ls *LoggerSet) Debug(v ...any) {
	if ls.Level > DebugLevel {
		return
	}
	ls.PrintlnAt(DebugLevel, v...)
}

// Sdebug returns the debug-level log-entries from all loggers in the default set.
func (ls *LoggerSet) Sdebug(v ...any) []string {
	if ls.Level > DebugLevel {
		return []string{}
	}
	return ls.SprintAt(DebugLevel, v...)
}

// Info prints an info-level log-entry to all loggers in the set.
func (ls *LoggerSet) Info(v ...any) {
	if ls.Level > InfoLevel {
		return
	}
	ls.PrintlnAt(InfoLevel, v...)
}

// Sinfo returns the info-level log-entries from all loggers in the default set.
func (ls *LoggerSet) Sinfo(v ...any) []string {
	if ls.Level > InfoLevel {
		return []string{}
	}
	return ls.SprintAt(InfoLevel, v...)
}

// Warning prints a warning-level log-entry to all loggers in the set.
func (ls *LoggerSet) Warning(v ...any) {
	if ls.Level > WarningLevel {
		return
	}
	ls.PrintlnAt(WarningLevel, v...)
}

// Swarning returns the warning-level log-entries from all loggers in the default set.
func (ls *LoggerSet) Swarning(v ...any) []string {
	if ls.Level > WarningLevel {
		return []string{}
	}
	return ls.SprintAt(WarningLevel, v...)
}

// Error prints a error-level log-entry to all loggers in the set.
func (ls *LoggerSet) Error(v ...any) {
	if ls.Level > ErrorLevel {
		return
	}
	ls.PrintlnAt(ErrorLevel, v...)
}

// Serror returns the error-level log-entries from all loggers in the default set.
func (ls *LoggerSet) Serror(v ...any) []string {
	if ls.Level > ErrorLevel {
		return []string{}
	}
	return ls.SprintAt(ErrorLevel, v...)
}

// Fatal prints an error-level log-entry to all loggers in the set and exits with 1.
func (ls *LoggerSet) Fatal(v ...any) {
	ls.PrintlnAt(ErrorLevel, v...)
	os.Exit(1)
}

// Panic prints an error-level log-entry to all loggers in the set and panics.
func (ls *LoggerSet) Panic(v ...any) {
	panic(ls.Serror(v...))
}

// SprintAt returns the formatted outputs of the default sets loggers at the given loglevel.
func SprintAt(level int, v ...any) []string {
	var outs []string
	for _, l := range DefaultLoggerSet.loggers {
		if level == l.Level() {
			outs = append(outs, l.SWrite(v...))
		}
	}
	return outs
}

// PrintAt prints with all the default sets loggers at the given loglevel.
func PrintAt(level int, v ...any) {
	for _, l := range DefaultLoggerSet.loggers {
		if level == l.Level() {
			l.Write(v...)
		}
	}
}

// PrintlnAt prints with all the default sets loggers at the given loglevel (and adds a newline).
func PrintlnAt(level int, v ...any) {
	DefaultLoggerSet.PrintAt(level, append(v, "\n")...)
}

// Println prints a level-less log-entry to all loggers in the default set (and adds a newline).
func Println(v ...any) {
	DefaultLoggerSet.PrintlnAt(NoLevel, v...)
}

// Print prints a level-less log-entry to all loggers in the default set.
func Print(v ...any) {
	DefaultLoggerSet.PrintAt(NoLevel, v...)
}

// Sprint returns the level-less log-entries from all loggers in the default set.
func Sprint(v ...any) {
	DefaultLoggerSet.SprintAt(NoLevel, v...)
}

// Debug prints a debug-level log-entry to all loggers in the default set.
func Debug(v ...any) {
	if DefaultLoggerSet.Level > DebugLevel {
		return
	}
	DefaultLoggerSet.PrintlnAt(DebugLevel, v...)
}

// Sdebug returns the debug-level log-entries from all loggers in the default set.
func Sdebug(v ...any) []string {
	if DefaultLoggerSet.Level > DebugLevel {
		return []string{}
	}
	return DefaultLoggerSet.SprintAt(DebugLevel, v...)
}

// Info prints an info-level log-entry to all loggers in the default set.
func Info(v ...any) {
	if DefaultLoggerSet.Level > InfoLevel {
		return
	}
	DefaultLoggerSet.PrintlnAt(InfoLevel, v...)
}

// Sinfo returns the info-level log-entries from all loggers in the default set.
func Sinfo(v ...any) []string {
	if DefaultLoggerSet.Level > InfoLevel {
		return []string{}
	}
	return DefaultLoggerSet.SprintAt(InfoLevel, v...)
}

// Warning prints a warning-level log-entry to all loggers in the default set.
func Warning(v ...any) {
	if DefaultLoggerSet.Level > WarningLevel {
		return
	}
	DefaultLoggerSet.PrintlnAt(WarningLevel, v...)
}

// Swarning returns the warning-level log-entries from all loggers in the default set.
func Swarning(v ...any) []string {
	if DefaultLoggerSet.Level > WarningLevel {
		return []string{}
	}
	return DefaultLoggerSet.SprintAt(WarningLevel, v...)
}

// Error prints a error-level log-entry to all loggers in the default set.
func Error(v ...any) {
	if DefaultLoggerSet.Level > ErrorLevel {
		return
	}
	DefaultLoggerSet.PrintlnAt(ErrorLevel, v...)
}

// Serror returns the error-level log-entries from all loggers in the default set.
func Serror(v ...any) []string {
	if DefaultLoggerSet.Level > ErrorLevel {
		return []string{}
	}
	return DefaultLoggerSet.SprintAt(ErrorLevel, v...)
}

// Fatal prints an error-level log-entry to all loggers in the default set and exits with 1.
func Fatal(v ...any) {
	DefaultLoggerSet.PrintlnAt(ErrorLevel, v...)
	os.Exit(1)
}

// Panic prints an error-level log-entry to all loggers in the default set and panics.
func Panic(v ...any) {
	panic(DefaultLoggerSet.Serror(v...))
}

/*

// SetOutput sets the loggers default output to the given writer
func (l *Logger) SetOutput(w io.Writer) {
	l.Output = w
}

// SetOutputFile sets the default loggers output to a file, it should be closed with Close()
// the given path will be appended to
func (l *Logger) SetOutputFile(path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.Output = f
	return nil
}

// Close closes the loggers output file if one is set
func (l *Logger) Close() error {
	f, isFile := l.Output.(*os.File)
	if !isFile {
		return nil
	}
	err := f.Close()
	return err
}

// sprint returns a log-entry at the given level with the given format, prefix, suffix and args
func sprint(loggerlevel int, loglevel int, format string, prefix string, suffix string, v ...any) string {
	if loggerlevel > loglevel {
		return ""
	}
	for range v {
		format = format + "%v "
	}
	format = format + "%s"
	return sprintf(loggerlevel, loglevel, format, prefix, suffix, v...)
}

// sprintf returns a log-entry at the given level with the given format, prefix, suffix and args
func sprintf(loggerlevel int, loglevel int, format string, prefix string, suffix string, v ...any) string {
	if loggerlevel > loglevel {
		return ""
	}
	p := append([]any{prefix}, v...)
	p = append(p, suffix)
	return fmt.Sprintf(format, p...)
}

// Println prints a level-less log-entry to the default writer
func Println(v ...any) {
	DefaultLogger.Write([]byte(DefaultLogger.Sprint(v...) + "\n"))
}

// Debug prints a log-entry at debug level to the default writer
func Debug(v ...any) {
	if DefaultLogger.Level > DebugLevel {
		return
	}
	DefaultLogger.Write([]byte(DefaultLogger.Sdebug(v...) + "\n"))
}

// Info prints a log-entry at info level to the default writer
func Info(v ...any) {
	if DefaultLogger.Level > InfoLevel {
		return
	}
	DefaultLogger.Write([]byte(DefaultLogger.Sinfo(v...) + "\n"))
}

// Warning prints a log-entry at warning level to the default writer
func Warning(v ...any) {
	if DefaultLogger.Level > WarningLevel {
		return
	}
	DefaultLogger.WriteErr([]byte(DefaultLogger.Swarning(v...) + "\n"))
}

// Error prints a log-entry at error level to the default writer
func Error(v ...any) {
	if DefaultLogger.Level > ErrorLevel {
		return
	}
	DefaultLogger.WriteErr([]byte(DefaultLogger.Serror(v...) + "\n"))
}

// Fatal prints a log-entry at the default error level and exits with 1
func Fatal(v ...any) {
	DefaultLogger.Fatal(v...)
}

// Panic prints a log-entry at the default error level and panics
func Panic(v ...any) {
	panic(DefaultLogger.Serror(v...))
}

// Println prints a level-less log-entry to the given writer
func (l *Logger) Println(v ...any) {
	l.Write([]byte(l.Sprint(v...) + "\n"))
}

// Sprint returns a level-less log-entry
func (l *Logger) Sprint(v ...any) string {
	return sprint(NoLevel, NoLevel, l.DefaultFormat, l.DefaultPrefix, l.DefaultSuffix, v...)
}

// Sprintf returns a level-less log-entry following the format
func (l *Logger) Sprintf(format string, v ...any) string {
	return sprintf(NoLevel, NoLevel, format, l.DefaultPrefix, l.DefaultSuffix, v...)
}

// Sdebugf returns a log-entry at debug level following the format
func (l *Logger) Sdebugf(format string, v ...any) string {
	return sprintf(l.Level, DebugLevel, format, l.DebugPrefix, l.DebugSuffix, v...)
}

// Sinfof returns a log-entry at info level following the format
func (l *Logger) Sinfof(format string, v ...any) string {
	return sprintf(l.Level, InfoLevel, format, l.InfoPrefix, l.InfoSuffix, v...)
}

// Swarningf returns a log-entry at warning level following the format
func (l *Logger) Swarningf(format string, v ...any) string {
	return sprintf(l.Level, WarningLevel, format, l.WarningPrefix, l.WarningSuffix, v...)
}

// Serrorf returns a log-entry at error level following the format
func (l *Logger) Serrorf(format string, v ...any) string {
	return sprintf(l.Level, ErrorLevel, format, l.ErrorPrefix, l.ErrorSuffix, v...)
}

// Fatal prints a log-entry at error level and exits with 1
func (l *Logger) Fatal(v ...any) {
	l.Error(v...)
	os.Exit(1)
}

// Panic prints a log-entry at error level and panics
func (l *Logger) Panic(v ...any) {
	panic(l.Serror(v...))
}

*/
