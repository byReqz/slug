// Package console defines a slug logger that outputs to the terminal i/o streams.
package console

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"time"
)

// ConsoleLogger defines a logger that outputs to console data streams (stdout and stderr).
type ConsoleLogger struct {
	level  int
	format string
	output io.Writer
}

// NewConsoleLogger returns a default ConsoleLogger that writes to stdout.
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		level:  0,
		format: "[" + time.Now().Format("2006/01/02 15:04:05") + "] %s\n",
		output: os.Stdout,
	}
}

// NewDefaultConsoleLoggers returns an array of ConsoleLoggers using the slug formatting. Warnings and errors will be printed to stderr, everything else to stdout.
func NewDefaultConsoleLoggers() []*ConsoleLogger {
	tn := time.Now().Format("2006/01/02 15:04:05")
	set := append([]*ConsoleLogger{},
		&ConsoleLogger{
			level:  -2,
			format: tn + " | %s",
			output: os.Stdout,
		},
		&ConsoleLogger{
			level:  -1,
			format: tn + " | " + color.MagentaString("Debug:") + " %s",
			output: os.Stdout,
		},
		&ConsoleLogger{
			level:  0,
			format: tn + " | " + color.CyanString("Info:") + " %s",
			output: os.Stdout,
		},
		&ConsoleLogger{
			level:  1,
			format: tn + " | " + color.YellowString("Warning:") + " %s",
			output: os.Stderr,
		},
		&ConsoleLogger{
			level:  2,
			format: tn + " | " + color.RedString("Error:") + " %s",
			output: os.Stderr,
		},
	)
	return set
}

// Level returns the loggers level.
func (cs *ConsoleLogger) Level() int {
	return cs.level
}

// SetLevel sets the loggers level.
func (cs *ConsoleLogger) SetLevel(lvl int) {
	cs.level = lvl
}

// Format returns the loggers format string.
func (cs *ConsoleLogger) Format() string {
	return cs.format
}

// SetFormat sets the loggers format string. The given input will always be condensed to a single string before the format is applied.
func (cs *ConsoleLogger) SetFormat(format string) {
	cs.format = format
}

// SWrite returns a string as formatted by the logger.
func (cs *ConsoleLogger) SWrite(data ...any) string {
	return fmt.Sprintf(cs.format, fmt.Sprint(data...))
}

// Write writes the given data to the loggers output.
func (cs *ConsoleLogger) Write(data ...any) {
	_ = cs.WriteE(data...)
}

// WriteE writes the given data to the loggers output and reports about errors.
func (cs *ConsoleLogger) WriteE(data ...any) error {
	_, err := fmt.Fprintf(cs.output, cs.format, fmt.Sprint(data...))
	return err
}