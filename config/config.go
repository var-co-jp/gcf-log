package config

import (
	"log"
	"time"
)

type LogType string

var (
	projectID = ""
)

const (
	DEBUG    = LogType("DEBUG")
	INFO     = LogType("INFO")
	WARN     = LogType("WARNING")
	ERROR    = LogType("ERROR")
	CRITICAL = LogType("CRITICAL")
	ALERT    = LogType("ALERT")
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
