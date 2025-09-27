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

	logger.Info("Executed successfully.")
}
