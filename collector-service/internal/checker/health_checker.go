package checker

import (
	"context"
	"net/http"
	"time"

	"github.com/Chintukr2004/collector/internal/models"
)

type Result struct {
	ServiceID   int
	ServiceName string
	URL         string
	Status      string
	Latency     time.Duration
	Error       error
}

func CheckHealth(service models.Service) Result {

	var lastErr error
	var latency time.Duration

	retries := 3

	for i := 0; i < retries; i++ {

		start := time.Now()

		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(service.Timeout)*time.Second,
		)

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, service.URL, nil)

		client := http.Client{}

		resp, err := client.Do(req)

		latency = time.Since(start)

		cancel()

		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {

			resp.Body.Close()

			return Result{
				ServiceID: service.ID,
				ServiceName: service.Name,
				URL:         service.URL,
				Status:      "UP",
				Latency:     latency,
			}
		}

		lastErr = err

		// delay before retry
		time.Sleep(1 * time.Second)
	}

	return Result{
		ServiceID: service.ID,
		ServiceName: service.Name,
		URL:         service.URL,
		Status:      "DOWN",
		Latency:     latency,
		Error:       lastErr,
	}
}
