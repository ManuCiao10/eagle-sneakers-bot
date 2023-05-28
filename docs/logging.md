# Logging

To implement logging in Golang, we can use the standard library's `log` package or considerer a third-party logging libraries. I will describe how to perform basic logging using `Logtail`, a public API of the tailscale logs service, with a a dashboard of the data included.

### Logging

The snippet code has been taken from the offcial repository of [eagle-sneakers-bot](../handler/logs/logtail.go#L65).

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
