package main

import (
	"os"
	"time"

	"go.uber.org/zap"
)

func main() {
	zapped, _ := zap.NewProduction()
	logger := zapped.Sugar()
	defer zapped.Sync()

	logger.Info("Starting application...")
	secondsToSleep := os.Getenv("SECONDS_TO_SLEEP")
	if secondsToSleep == "" {
		secondsToSleep = "2"
	}

	secondsToSleepDuration, err := time.ParseDuration(secondsToSleep + "s")
	if err != nil {
		logger.Errorf("Invalid SECONDS_TO_SLEEP value: %v", secondsToSleep, zap.Error(err))
		return
	}

	logger.Infof("Sleeping for %s seconds...", secondsToSleep)
	time.Sleep(secondsToSleepDuration)
	logger.Info("Sleep finished.")

	if os.Getenv("SHOULD_FAIL") == "true" {
		logger.Error("Exiting with failure as per SHOULD_FAIL=true")
		panic("Exiting with failure as per SHOULD_FAIL=true")
	}

	if os.Getenv("SHOULD_FAIL_RANDOMLY") == "true" {
		if time.Now().UnixNano()%2 == 0 {
			// Simple pseudo-random check using current time's nanoseconds (no extra import needed)
			logger.Error("Panicking as per SHOULD_FAIL_RANDOMLY=true (randomly chosen)")
			panic("Exiting with random failure as per SHOULD_FAIL_RANDOMLY=true")
		}
		logger.Info("Random check passed; not panicking.")
	}

	logger.Info("Executed successfully.")
}
