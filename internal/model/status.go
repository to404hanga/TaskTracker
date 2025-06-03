package model

import "strings"

type StatusCode int

const (
	Unknown StatusCode = iota
	Todo
	InProgress
	Done
)

func ToString(s StatusCode) string {
	switch s {
	case Todo:
		return "todo"
	case InProgress:
		return "in-progress"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}

func FromString(s string) StatusCode {
	switch strings.ToLower(s) {
	case "todo":
		return Todo
	case "in-progress":
		return InProgress
	case "done":
		return Done
	default:
		return Unknown
	}
}
