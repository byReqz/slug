//nolint
package main

import (
	"fmt"
	"github.com/byReqz/slug"
)

func logrunner(l *slug.Logger) {
	l.Println("test123", 22292, true)
	l.Debug("test123", 22292, true)
	l.Info("test123", 22292, true)
	l.Warning("test123", 22292, true)
	l.Error("test123", 22292, true)
}

func main() {
	l := slug.NewLogger()
	l.SetLevel(slug.DebugLevel)
	err := l.SetOutputFile("demo.log")
	if err != nil {
		slug.Fatal(err)
	}
	defer l.Close()
	l.Println("log with colors")
	logrunner(l)

	l.DisableColor()
	l.Write([]byte("\n"))
	l.Println("log without colors")
	logrunner(l)

	fmt.Println("output at ./demo.log")
}
