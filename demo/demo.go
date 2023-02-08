// nolint
package main

import (
	"fmt"
	"github.com/byReqz/slug"
)

func logrunner(l *slug.LoggerSet) {
	l.Println("test123", 22292, true)
	l.Debug("test123", 22292, true)
	l.Info("test123", 22292, true)
	l.Warning("test123", 22292, true)
	l.Error("test123", 22292, true)
}

func main() {
	l := slug.NewConsoleLoggerSet()

	l.Level = slug.Disabled
	fmt.Println("disabled:")
	logrunner(l)

	l.Level = slug.NoLevel
	fmt.Println("at no level:")
	logrunner(l)

	l.Level = slug.DebugLevel
	fmt.Println("at debug level:")
	logrunner(l)

	l.Level = slug.InfoLevel
	fmt.Println("at info level:")
	logrunner(l)

	l.Level = slug.WarningLevel
	fmt.Println("at warning level:")
	logrunner(l)

	l.Level = slug.ErrorLevel
	fmt.Println("at error level:")
	logrunner(l)
}
