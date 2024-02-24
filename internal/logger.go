package internal

import (
	"time"

	"github.com/fatih/color"
)

type LogSeverity int

const (
	Info LogSeverity = iota
	Warning
	Error
)

type LogMessage struct {
	Message  string
	Severity LogSeverity
	Time     time.Time
}

var logs = []LogMessage{}

var (
	infoColor    = color.New(color.FgWhite)
	warningColor = color.New(color.FgYellow)
	errorColor   = color.New(color.FgRed)
)

func (newLog *LogMessage) LogMessage() {
	switch newLog.Severity {
	case Info:
		infoColor.Printf("%v", newLog)
	case Warning:
		warningColor.Printf("%v", newLog)
	case Error:
		errorColor.Printf("%v", newLog)
	}
}

func CreateLog(msg string, sev LogSeverity) LogMessage {
	l := LogMessage{
		Message:  msg,
		Severity: sev,
		Time:     time.Now(),
	}

	logs = append(logs, l)
	l.LogMessage()
	return l
}
