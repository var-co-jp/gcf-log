# ✍️ GCF-LOG


This is the easiest library for GCF logging.

## Install


```bash
go get -u "github.com/var-co-jp/gcf-log"
```

## Usage

### HTTP Trigger

```go
import "github.com/var-co-jp/gcf-log"

var projectID = "your GCP project ID here!"

func HttpTrigger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// use Init() with args context, GCP projectID, and **http.Request** 
	ctx = gcflog.Init(ctx, projectID, r)
	gcflog.Info(ctx, "Log started.")
	gcflog.Warning(ctx, "This is warning")
}
```

### **Event Trigger**

```go
import "github.com/var-co-jp/gcf-log"

var projectID = "your GCP project ID here!"

func HttpTrigger(ctx context.Context, m PubSubMessage) {
	// use Init() with args context, GCP projectID, and nil ****
	ctx = gcflog.Init(ctx, projectID, nil)
	gcflog.Debug(ctx, "Log started.")
	gcflog.Error(ctx, "This is error")
}
```

This library can use any levels of **Severity**.

For more information, see see the URL below.

[https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity](https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity)

## Licence

MIT
