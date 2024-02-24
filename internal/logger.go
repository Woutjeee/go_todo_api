package internal

import (
	"encoding/json"
	"fmt"
	"os"
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
	Message  string      `json:"message"`
	Severity LogSeverity `json:"severity"`
	Time     time.Time   `json:"time"`
}

type JsonLogMessage struct {
	Logs map[string][]LogMessage
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

func StartLogger() {
	currentTime := time.Now()
	formattedDate := currentTime.Format("02-01-2006")
	getCurrentLogs(formattedDate)
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

func getCurrentLogs(date string) ([]LogMessage, error) {
	filename := "logs.json"

	// Get current logs by date.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		CreateLog("Logs file does not exists, creating it now", Info)

		file, err := os.Create(filename)
		if err != nil {
			msg := fmt.Sprintf("Error creating file: %v", err)
			CreateLog(msg, Error)
			return nil, err
		}

		defer file.Close()
		CreateLog("Logs file created", Info)
	} else if err != nil {
		msg := fmt.Sprintf("Error checking file: %v", err)
		CreateLog(msg, Error)
	} else {
		file, err := os.Open(filename)
		if err != nil {
			msg := fmt.Sprintf("Error opening file: %v", err)
			CreateLog(msg, Error)
			return nil, err
		}

		defer file.Close()

		var data JsonLogMessage
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			msg := fmt.Sprintf("Error decoding JSON: %v", err)
			CreateLog(msg, Error)
			return nil, err
		}

		// Loop through data here.
		fmt.Println(data)
		logs, found := data.Logs[date]
		if !found {
			fmt.Printf("No logs found for date: %s\n", date)
			return nil, nil // Retuirn WARNING here that no logs have been made yet. (Or atleast written to the json file.)
		}

		for _, log := range logs {
			fmt.Printf("Message: %s, Severity: %s\n", log.Message, log.Severity)
		}
	}

	return nil, nil
}
