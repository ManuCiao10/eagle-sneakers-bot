# Logging

To implement logging in Golang, we can use the standard library's `log` package or consider third-party logging libraries. In this example, I will describe how to perform basic logging using `Logtail`, a public API of the Tailscale logs service, which includes a dashboard for viewing the logged data.

### Logging

The code snippet has been taken from the official repository of [eagle-sneakers-bot](../handler/logs/logtail.go#L65).

```go
func flushLogs() {
	logsMutex.RLock()

	if len(Logs) > 0 {
		requestBody := getRequestBodyForFlush()
		logsMutex.RUnlock()

		client, err := client.NewClient()
		if err != nil {
			return
		}

		_, err = client.NewRequest().
			SetMethod("POST").
			SetURL("https://in.logtail.com").
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "*/*").
			SetHeader("Authorization", "Bearer authkey").
			SetJSONBody(requestBody).
			Do()

		if err != nil {
			return
		} else {
			logsMutex.Lock()
			Logs = make(map[LogtailData]int)
			logsMutex.Unlock()
		}
	} else {
		logsMutex.RUnlock()
	}
}
```
