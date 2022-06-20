package main

import (
	"github.com/byReqz/slug"
)

func main() {
	l := slug.NewLogger()
	l.Level = -1
	l.Debug("test123", 22292, true)
}
