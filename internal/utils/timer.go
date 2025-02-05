package utils

import "time"

func Timer(interval time.Duration, f func()) {
	InfoLogger.Printf("Starting timer with interval: %v", interval)
	f()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		f()
	}

}
