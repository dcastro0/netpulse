package monitor

import (
	"net/http"
	"time"
)

type Result struct {
	URL      string
	Status   int
	Duration time.Duration
	Err      error
}

func Check(url string, timeout int) Result {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		return Result{URL: url, Err: err, Duration: duration}
	}
	defer resp.Body.Close()

	return Result{
		URL:      url,
		Status:   resp.StatusCode,
		Duration: duration,
	}
}	