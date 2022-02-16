package main

import (
	"golang-logs/loggers"
)

func main() {
	loggers.LogrusLogger()

	loggers.LogrusError("Error parsing message")
	// loggers.LogrusErrorWithMetadata("Failed to process request", log.Fields{
	// 	"document": "6181681681",
	// 	"error":    "Error calling API",
	// })
}
