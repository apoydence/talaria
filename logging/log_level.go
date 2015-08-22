package logging

import (
	"bytes"
	"fmt"
)

type LogLevel int

const (
	CRITICAL LogLevel = iota
	ERROR
	INFO
	DEBUG
)

func (l LogLevel) String() string {
	switch l {
	case CRITICAL:
		return "CRITICAL"
	case ERROR:
		return "ERROR"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		return "INVALID"
	}
}

func (l LogLevel) MarshalJSON() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *LogLevel) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, `"`)
	strData := string(data)
	switch strData {
	case "CRITICAL":
		*l = CRITICAL
	case "ERROR":
		*l = ERROR
	case "INFO":
		*l = INFO
	case "DEBUG":
		*l = DEBUG
	default:
		return fmt.Errorf("Unknown LogLevel: %v", strData)
	}
	return nil
}