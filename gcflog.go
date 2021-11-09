package gcflog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/var-co-jp/gcf-log/config"
	"go.opencensus.io/trace"
)

type key int

var header key

func Init(original context.Context, projectID string, r *http.Request) (ctx context.Context) {
	config.SetProjectID(projectID)
	if r != nil {
		headerValue := fetchHeader(r)
		ctx = context.WithValue(original, header, headerValue)
		return
	}
	pctx := trace.FromContext(original).SpanContext()
	headerValue := generateHeader(pctx)
	ctx = context.WithValue(original, header, headerValue)
	return
}

func fetchHeader(r *http.Request) (header string) {
	header = r.Header.Get("X-Cloud-Trace-Context")
	if header == "" {
		header = "00000000000000000000000000000000/0000000000000000;o=TRACE_TRUE"
	}
	return
}

func generateHeader(pctx trace.SpanContext) (header string) {
	header = fmt.Sprintf("%s/%s;o=TRACE_TRUE", pctx.TraceID.String(), pctx.SpanID.String())
	return
}

func getSourceLocation() (srcJsonFmt map[string]string) {
	pc, file, line, _ := runtime.Caller(2)
	fname := filepath.Base(file)
	funcName := runtime.FuncForPC(pc).Name()
	srcJsonFmt = map[string]string{
		"file":     fname,
		"line":     strconv.Itoa(line),
		"function": funcName,
	}
	return
}

func getTrace(ctx context.Context) (trace string) {
	header := ctx.Value(header).(string)
	traceContext := strings.Split(header, "/")
	projectID := config.GetProjectID()
	if len(traceContext) < 2 {
		return fmt.Sprintf("projects/%s/traces/00000000000000000000000000000000", projectID)
	}
	traceId := traceContext[0]
	if traceId == "" {
		return fmt.Sprintf("projects/%s/traces/00000000000000000000000000000000", projectID)
	}
	trace = fmt.Sprintf("projects/%s/traces/%s", projectID, traceId)
	return
}

func getSpan(ctx context.Context) (span string) {
	header := ctx.Value(header).(string)
	traceContext := strings.Split(header, "/")
	if len(traceContext) < 2 {
		return "0000000000000000"
	}
	spanId := strings.Split(traceContext[1], ";")[0]
	if spanId == "" {
		return "0000000000000000"
	}
	span = spanId
	return
}

func SetLogLevel(level string) {
	config.SetLogLevel(level)
}

func Debug(ctx context.Context, message string) {
	if config.DEBUG.Level <= config.GetLogLevel() {
		return
	}
	logDebug := config.LogFormat{
		Severity:       config.DEBUG.LogType,
		Message:        message,
		Time:           time.Now(),
		SourceLocation: getSourceLocation(),
		TraceId:        getTrace(ctx),
		SpanId:         getSpan(ctx),
		TraceSample:    false,
	}
	bytes, _ := json.Marshal(logDebug)
	fmt.Println(string(bytes))
}

func Info(ctx context.Context, message string) {
	if config.INFO.Level <= config.GetLogLevel() {
		return
	}
	logInfo := config.LogFormat{
		Severity:       config.INFO.LogType,
		Message:        message,
		Time:           time.Now(),
		SourceLocation: getSourceLocation(),
		TraceId:        getTrace(ctx),
		SpanId:         getSpan(ctx),
		TraceSample:    false,
	}
	bytes, _ := json.Marshal(logInfo)
	fmt.Println(string(bytes))
}

func Warn(ctx context.Context, message string) {
	if config.WARN.Level <= config.GetLogLevel() {
		return
	}
	logWarn := config.LogFormat{
		Severity:       config.WARN.LogType,
		Message:        message,
		Time:           time.Now(),
		SourceLocation: getSourceLocation(),
		TraceId:        getTrace(ctx),
		SpanId:         getSpan(ctx),
		TraceSample:    false,
	}
	bytes, _ := json.Marshal(logWarn)
	fmt.Println(string(bytes))
}

func Error(ctx context.Context, message string) {
	if config.ERROR.Level <= config.GetLogLevel() {
		return
	}
	logError := config.LogFormat{
		Severity:       config.ERROR.LogType,
		Message:        message,
		Time:           time.Now(),
		SourceLocation: getSourceLocation(),
		TraceId:        getTrace(ctx),
		SpanId:         getSpan(ctx),
		TraceSample:    false,
	}
	bytes, _ := json.Marshal(logError)
	fmt.Println(string(bytes))
}

func Critical(ctx context.Context, message string) {
	if config.CRITICAL.Level <= config.GetLogLevel() {
		return
	}
	logCritical := config.LogFormat{
		Severity:       config.CRITICAL.LogType,
		Message:        message,
		Time:           time.Now(),
		SourceLocation: getSourceLocation(),
		TraceId:        getTrace(ctx),
		SpanId:         getSpan(ctx),
		TraceSample:    false,
	}
	bytes, _ := json.Marshal(logCritical)
	fmt.Println(string(bytes))
}

func Alert(ctx context.Context, message string) {
	if config.ALERT.Level <= config.GetLogLevel() {
		return
	}
	logAlert := config.LogFormat{
		Severity:       config.ALERT.LogType,
		Message:        message,
		Time:           time.Now(),
		SourceLocation: getSourceLocation(),
		TraceId:        getTrace(ctx),
		SpanId:         getSpan(ctx),
		TraceSample:    false,
	}
	bytes, _ := json.Marshal(logAlert)
	fmt.Println(string(bytes))
}
