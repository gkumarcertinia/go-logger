package go-logger

import "log"

func logInfo(message string) {
	log.Printf("INFO - %v ", message)
}

func logWarning(message string) {
	log.Printf("WARN - %v ", message)
}

func logError(message string) {
	log.Printf("ERROR - %v ", message)
}

func logDebug(message string) {
	log.Printf("DEBUG - %v ", message)
}
