package msl

import (
	"fmt"
)

type Event struct {
	Type     string
	Pattern  string
	Location string
	Level    string
	Line     int

	Command Command
}

func (ev *Event) Name() string {
	return fmt.Sprintf("ev_%s_%dl", ev.Type, ev.Line)
}

func (ev *Event) Flags() string {
	return "--level " + ev.Level + " --location " + ev.Location
}

func (ev *Event) Match() string {
	if ev.Type == "join" {
		return "^(?P<user>[^\\s@!]+)[^\\s]*\\s+JOIN\\s+"
	}
	return ev.Pattern
}
