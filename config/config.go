package config

import (
	"log"
	"time"
)

type LogType string

type Log struct {
	LogType LogType
	Level   int
}

var (
	projectID = ""
	logLevel  = DEBUG
)

var (
	DEBUG    = Log{LogType("DEBUG"), 0}
	INFO     = Log{LogType("INFO"), 1}
	WARN     = Log{LogType("WARN"), 2}
	ERROR    = Log{LogType("ERROR"), 3}
	CRITICAL = Log{LogType("CRITICAL"), 4}
	ALERT    = Log{LogType("ALERT"), 5}
)

type LogFormat struct {
	Severity       LogType           `json:"severity"`
	Message        string            `json:"message"`
	Time           time.Time         `json:"time"`
	SourceLocation map[string]string `json:"logging.googleapis.com/sourceLocation"`
	TraceId        string            `json:"logging.googleapis.com/trace"`
	SpanId         string            `json:"logging.googleapis.com/spanId"`
	TraceSample    bool              `json:"logging.googleapis.com/trace_sampled"`
}

func SetProjectID(_projectID string) {
	if _projectID == "" {
		log.Printf("WARN: Project ID value is empty")
	}
	projectID = _projectID
}

func GetProjectID() string {
	return projectID
}

func SetLogLevel(level string) {
	switch level {
	case "DEBUG":
		logLevel = DEBUG
	case "INFO":
		logLevel = INFO
	case "WARN":
		logLevel = WARN
	case "ERROR":
		logLevel = ERROR
	case "CRITICAL":
		logLevel = CRITICAL
	case "ALERT":
		logLevel = ALERT
	default:
		log.Printf("WARN: Log Level value is invalid")
	}
}

func GetLogLevel() int {
	return logLevel.Level
}
