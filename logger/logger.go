package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	statusInfo  = "Info"
	statusError = "Error"
)

// LogMsg ...
type LogMsg struct {
	Status      string
	Time        time.Time
	Metadata    string
	Description string
}

// SendInfo ...
func SendInfo(md string, desc string) {

	lm := &LogMsg{
		Status:      statusInfo,
		Time:        time.Now(),
		Metadata:    md,
		Description: desc,
	}

	b, err := json.Marshal(lm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Fprintln(os.Stdout, string(b))
}

// SendError ...
func SendError(md string, desc string) {

	lm := &LogMsg{
		Status:      statusError,
		Time:        time.Now(),
		Metadata:    md,
		Description: desc,
	}

	b, err := json.Marshal(lm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Fprintln(os.Stderr, string(b))
}
