package slug

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"time"
)

var DefaultLogger *Logger // logger using default configuration

var (
	NoLevel      = -2 // log without level, level-less logs will always be printed regardless of logger level
	DebugLevel   = -1 // log at >= debug level
	InfoLevel    = 0  // log at >= info level
	WarningLevel = 1  // log at >= warning level
	ErrorLevel   = 2  // log at >= error level
	Disabled     = 3  // log at > error level (disable the logger)
)

type Logger struct {
	Level         int       // the minimum log level the logger will output
	Output        io.Writer // output of the logger
	DefaultFormat string    // formatstring for level-less logs
	DefaultPrefix string    // message prefix for level-less logs
	DefaultSuffix string    // message suffix for level-less logs
	DebugFormat   string    // formatstring for debug level logs
	DebugPrefix   string    // message prefix for debug level logs
	DebugSuffix   string    // message suffix for debug level logs
	InfoFormat    string    // formatstring for info level logs
	InfoPrefix    string    // message prefix for info level logs
	InfoSuffix    string    // message suffix for info level logs
	WarningFormat string    // formatstring for warning level logs
	WarningPrefix string    // message prefix for warning level logs
	WarningSuffix string    // message suffix for warning level logs
	ErrorFormat   string    // formatstring for error level logs
	ErrorPrefix   string    // message prefix for error level logs
	ErrorSuffix   string    // message suffix for error level logs
}

func init() {
	DefaultLogger = NewLogger() // set DefaultLogger
}

// NewLogger returns a new default Logger and should be used in most cases
func NewLogger() *Logger {
	var l Logger
	l.Output = os.Stdout
	l.DefaultFormat = "%s"
	l.DefaultPrefix = fmt.Sprint(time.Now().Format("2006/01/02 15:04:05")) + " | "
	l.DebugFormat = l.DefaultFormat
	l.DebugPrefix = l.DefaultPrefix + color.MagentaString("Debug: ")
	l.DebugSuffix = l.DefaultSuffix
	l.InfoFormat = l.DefaultFormat
	l.InfoPrefix = l.DefaultPrefix + color.CyanString("Info: ")
	l.InfoSuffix = l.DefaultSuffix
	l.WarningFormat = l.DefaultFormat
	l.WarningPrefix = l.DefaultPrefix + color.YellowString("Warning: ")
	l.WarningSuffix = l.DefaultSuffix
	l.ErrorFormat = l.DefaultFormat
	l.ErrorPrefix = l.DefaultPrefix + color.RedString("Error: ")
	l.ErrorSuffix = l.DefaultSuffix
	return &l
}

// SetOutput sets the loggers output to the given writer
func (l *Logger) SetOutput(w io.Writer) {
	l.Output = w
}

// SetOutputFile sets the loggers output to a file, it should be closed with Close()
// the given path will be appended to
func (l *Logger) SetOutputFile(path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.Output = f
	return nil
}

// SetLevel sets the loggers level
func (l *Logger) SetLevel(lvl int) {
	l.Level = lvl
}

// DisableColor removes the coloring from the default level prefixes/suffixes
func (l *Logger) DisableColor() {
	l.DebugPrefix = l.DefaultPrefix + "Debug: "
	l.InfoPrefix = l.DefaultPrefix + "Info: "
	l.WarningPrefix = l.DefaultPrefix + "Warning: "
	l.ErrorPrefix = l.DefaultPrefix + "Error: "
}

// EnableColor resets the coloring on the default level prefixes/suffixes
func (l *Logger) EnableColor() {
	l.DebugPrefix = l.DefaultPrefix + color.MagentaString("Debug: ")
	l.InfoPrefix = l.DefaultPrefix + color.CyanString("Info: ")
	l.WarningPrefix = l.DefaultPrefix + color.YellowString("Warning: ")
	l.ErrorPrefix = l.DefaultPrefix + color.RedString("Error: ")
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

// sprintln returns a log-entry at the given level with the given format, prefix, suffix and args + newline at the end
func sprintln(loggerlevel int, loglevel int, format string, prefix string, suffix string, v ...any) string {
	if loggerlevel > loglevel {
		return ""
	}
	for range v {
		format = format + "%v "
	}
	format = format + "%s"
	return sprintf(loggerlevel, loglevel, format, prefix, suffix, v...) + "\n"
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
	DefaultLogger.Write([]byte(DefaultLogger.Sprintln(v...)))
}

// Debug prints a log-entry at debug level to the default writer
func Debug(v ...any) {
	if DefaultLogger.Level > DebugLevel {
		return
	}
	DefaultLogger.Write([]byte(DefaultLogger.Sdebugln(v...)))
}

// Info prints a log-entry at info level to the default writer
func Info(v ...any) {
	if DefaultLogger.Level > InfoLevel {
		return
	}
	DefaultLogger.Write([]byte(DefaultLogger.Sinfoln(v...)))
}

// Warning prints a log-entry at warning level to the default writer
func Warning(v ...any) {
	if DefaultLogger.Level > WarningLevel {
		return
	}
	DefaultLogger.Write([]byte(DefaultLogger.Swarningln(v...)))
}

// Error prints a log-entry at error level to the default writer
func Error(v ...any) {
	if DefaultLogger.Level > ErrorLevel {
		return
	}
	DefaultLogger.Write([]byte(DefaultLogger.Serrorln(v...)))
}

// Fatal prints a log-entry at the default error level and exits with 1
func Fatal(v ...any) {
	DefaultLogger.Fatal(v...)
}

// Println prints a level-less log-entry to the given writer
func (l *Logger) Println(v ...any) {
	l.Write([]byte(l.Sprintln(v...)))
}

// Sprintln returns a level-less log-entry with a newline at the end
func (l *Logger) Sprintln(v ...any) string {
	return sprintln(NoLevel, NoLevel, l.DefaultFormat, l.DefaultPrefix, l.DefaultSuffix, v...)
}

// Sprintf returns a level-less log-entry following the format
func (l *Logger) Sprintf(format string, v ...any) string {
	return sprintf(NoLevel, NoLevel, format, l.DefaultPrefix, l.DefaultSuffix, v...)
}

// Debug prints a log-entry at debug level to the given writer
func (l *Logger) Debug(v ...any) {
	if l.Level > DebugLevel {
		return
	}
	l.Write([]byte(l.Sdebugln(v...)))
}

// Sdebugln returns a log-entry at debug level with a newline at the end
func (l *Logger) Sdebugln(v ...any) string {
	return sprintln(l.Level, DebugLevel, l.DebugFormat, l.DebugPrefix, l.DebugSuffix, v...)
}

// Sdebugf returns a log-entry at debug level following the format
func (l *Logger) Sdebugf(format string, v ...any) string {
	return sprintf(l.Level, DebugLevel, format, l.DebugPrefix, l.DebugSuffix, v...)
}

// Info prints a log-entry at info level to the given writer
func (l *Logger) Info(v ...any) {
	if l.Level > InfoLevel {
		return
	}
	l.Write([]byte(l.Sinfoln(v...)))
}

// Sinfoln returns a log-entry at info level with a newline at the end
func (l *Logger) Sinfoln(v ...any) string {
	return sprintln(l.Level, InfoLevel, l.InfoFormat, l.InfoPrefix, l.InfoSuffix, v...)
}

// Sinfof returns a log-entry at info level following the format
func (l *Logger) Sinfof(format string, v ...any) string {
	return sprintf(l.Level, InfoLevel, format, l.InfoPrefix, l.InfoSuffix, v...)
}

// Warning prints a log-entry at warning level to the given writer
func (l *Logger) Warning(v ...any) {
	if l.Level > WarningLevel {
		return
	}
	l.Write([]byte(l.Swarningln(v...)))
}

// Swarningln returns a log-entry at warning level with a newline at the end
func (l *Logger) Swarningln(v ...any) string {
	return sprintln(l.Level, WarningLevel, l.WarningFormat, l.WarningPrefix, l.WarningSuffix, v...)
}

// Swarningf returns a log-entry at warning level following the format
func (l *Logger) Swarningf(format string, v ...any) string {
	return sprintf(l.Level, WarningLevel, format, l.WarningPrefix, l.WarningSuffix, v...)
}

// Error prints a log-entry at error level to the given writer
func (l *Logger) Error(v ...any) {
	if l.Level > ErrorLevel {
		return
	}
	l.Write([]byte(l.Serrorln(v...)))
}

// Serrorln returns a log-entry at error level with a newline at the end
func (l *Logger) Serrorln(v ...any) string {
	return sprintln(l.Level, ErrorLevel, l.ErrorFormat, l.ErrorPrefix, l.ErrorSuffix, v...)
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

// Write bytes to the logger's writer, falls back to stdout if error and panics if even that fails
func (l *Logger) Write(b []byte) {
	_, err := l.Output.Write(b)
	if err != nil {
		_, err = os.Stdout.Write(append([]byte("Failed printing to given output writer, falling back to stdout: "+err.Error()+"\n"), b...)) // try to fall back to stdout if error occurred
		if err != nil {
			panic("Failed printing to given output writer and also failed falling back to stdout")
		}
	}
}
