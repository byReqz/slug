package slug

import (
	"fmt"
	"github.com/fatih/color"
)

var DefaultLogger Logger

var (
	DebugLevel   = -1
	InfoLevel    = 0
	WarningLevel = 1
	ErrorLevel   = 2
)

type Logger struct {
	Enabled       bool
	Color         bool
	Level         int
	DebugFormat   string
	DebugPrefix   string
	InfoFormat    string
	InfoPrefix    string
	WarningFormat string
	WarningPrefix string
	ErrorFormat   string
	ErrorPrefix   string
}

func NewLogger() *Logger {
	var l Logger
	l.Enabled = true
	l.Color = true
	l.DebugFormat = "%s "
	l.DebugPrefix = color.CyanString("Debug:")
	l.InfoFormat = "%s "
	l.InfoPrefix = color.MagentaString("Info:")
	l.WarningFormat = "%s "
	l.WarningPrefix = color.YellowString("Warning:")
	l.ErrorFormat = "%s "
	l.ErrorPrefix = color.RedString("Error:")
	return &l
}

func (l *Logger) Println(v ...any) {

}

func (l *Logger) Sprintln(v ...any) string {
	return ""
}

func (l *Logger) Debug(v ...any) {
	if l.Level > -1 {
		return
	}
	fmt.Print(l.SDebugln(v...))
}

func (l *Logger) SDebugln(v ...any) string {
	if l.Level > -1 {
		return ""
	}
	format := l.DebugFormat
	for _, _ = range v {
		format = format + "%v "
	}
	return l.SDebugf(format, v...) + "\n"
}

func (l *Logger) SDebugf(format string, v ...any) string {
	if l.Level > -1 {
		return ""
	}
	p := append([]any{l.DebugPrefix}, v...)
	return fmt.Sprintf(format, p...)
}
