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
	DebugLevel   = -1
	InfoLevel    = 0
	WarningLevel = 1
	ErrorLevel   = 2
)

type Logger struct {
	Enabled       bool      // state of logger
	Color         bool      // whether to use colors
	Level         int       // the minimum log level the logger will output
	Output        io.Writer // output of the logger
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
	l.Enabled = true
	l.Color = true
	l.Output = os.Stdout
	l.DebugFormat = "%s "
	l.DebugPrefix = color.MagentaString("Debug:")
	l.DebugSuffix = "| " + fmt.Sprint(time.Now().Round(time.Second))
	l.InfoFormat = "%s "
	l.InfoPrefix = color.CyanString("Info:")
	l.InfoSuffix = "| " + fmt.Sprint(time.Now().Round(time.Second))
	l.WarningFormat = "%s "
	l.WarningPrefix = color.YellowString("Warning:")
	l.WarningSuffix = "| " + fmt.Sprint(time.Now().Round(time.Second))
	l.ErrorFormat = "%s "
	l.ErrorPrefix = color.RedString("Error:")
	l.ErrorSuffix = "| " + fmt.Sprint(time.Now().Round(time.Second))
	return &l
}

func (l Logger) Println(v ...any) {

}

func (l Logger) Sprintln(v ...any) string {
	return ""
}

// Debug prints a log-entry at debug level to the given writer
func (l Logger) Debug(v ...any) {
	if l.Level > -1 {
		return
	}
	l.Write([]byte(l.SDebugln(v...)))
}

// SDebugln returns a log-entry at debug level with a newline at the end
func (l Logger) SDebugln(v ...any) string {
	if l.Level > -1 {
		return ""
	}
	for range v {
		l.DebugFormat = l.DebugFormat + "%v "
	}
	l.DebugFormat = l.DebugFormat + "%s"
	return l.SDebugf(l.DebugFormat, v...) + "\n"
}

// SDebugf returns a log-entry at debug level following the format f
func (l Logger) SDebugf(format string, v ...any) string {
	if l.Level > -1 {
		return ""
	}
	p := append([]any{l.DebugPrefix}, v...)
	p = append(p, l.DebugSuffix)
	return fmt.Sprintf(format, p...)
}

// Error prints a log-entry at error level
func (l Logger) Error(v ...any) {
	if l.Level > -1 {
		return
	}
	l.Write([]byte(l.SErrorln(v...)))
}

// SErrorln returns a log-entry at error level with a newline at the end
func (l Logger) SErrorln(v ...any) string {
	if l.Level > -1 {
		return ""
	}
	for range v {
		l.ErrorFormat = l.ErrorFormat + "%v "
	}
	l.ErrorFormat = l.ErrorFormat + "%s"
	return l.SErrorf(l.ErrorFormat, v...) + "\n"
}

// SErrorf returns a log-entry at error level following the format f
func (l Logger) SErrorf(format string, v ...any) string {
	if l.Level > -1 {
		return ""
	}
	p := append([]any{l.ErrorPrefix}, v...)
	p = append(p, l.ErrorSuffix)
	return fmt.Sprintf(format, p...)
}

// Fatal prints a log-entry at error level and exits with 1
func (l Logger) Fatal(v ...any) {
	l.Error(v...)
	os.Exit(1)
}

// Write bytes to the loggers writer
func (l Logger) Write(b []byte) {
	_, err := l.Output.Write(b)
	if err != nil {
		_, err = os.Stdout.Write(append([]byte("Failed printing to given output writer, falling back to stdout: "+err.Error()+"\n"), b...)) // try to fall back to stdout if error occured
		if err != nil {
			panic("Failed printing to given output writer and also failed falling back to stdout")
		}
	}
}
